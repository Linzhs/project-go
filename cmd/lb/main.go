package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"project-go/internal/lb"
	"strings"
	"time"
)

// There is a lot we can do to improve our tiny load balancer.
//
// For example,
//
// Use a heap for sort out alive backends to reduce search surface
// Collect statistics
// Implement weighted round-robin/least connections
// Add support for a configuration file

var serverPool lb.ServerPool

var (
	serverList string
	port       int
)

func init() {
	flag.StringVar(&serverList, "backends", "", "Load balanced backends, ue commas to separate")
	flag.IntVar(&port, "port", 8080, "Port to serve")
	flag.Parse()
}

func doLoadBalance(w http.ResponseWriter, r *http.Request) {
	attempts := lb.GetRetryFromContext(r)
	if attempts > 3 {
		log.Printf("%s(%s) Max attempts reached, terminating \n", r.RemoteAddr, r.URL.Path)
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}

	peer := serverPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func healthCheck() {
	t := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-t.C:
			log.Println("Starting health check...")
			serverPool.HealthCheck() // do real health check
			log.Println("Health check completed")
		}
	}
}

func main() {

	if len(serverList) == 0 {
		log.Fatal("Please provide one or more backends to load balance")
	}

	// parse servers
	for _, t := range strings.Split(serverList, ",") {
		serverUrl, err := url.Parse(t)
		if err != nil {
			log.Fatal(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(serverUrl)
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
			log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
			if retries := lb.GetRetryFromContext(request); retries < 3 {
				select {
				case <-time.After(10 * time.Millisecond):
					ctx := context.WithValue(request.Context(), lb.Retry, retries+1)
					proxy.ServeHTTP(writer, request.WithContext(ctx))
				}
				return
			}

			serverPool.MarkBackendStatus(serverUrl, false)

			attempts := lb.GetAttemptsFromContext(request)
			log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)
			ctx := context.WithValue(request.Context(), attempts, attempts+1)
			doLoadBalance(writer, request.WithContext(ctx))
		}

		serverPool.AddBackend(&lb.Backend{
			Url:          serverUrl,
			Alive:        true,
			ReverseProxy: proxy,
		})
		log.Printf("Configured server: %s\n", serverUrl)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(doLoadBalance),
	}

	go healthCheck()

	log.Printf("Load balancer started at: %d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
