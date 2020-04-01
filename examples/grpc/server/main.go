package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project-go/pkg/grpclb"
	"project-go/pkg/rpc/pb/helloworld"
	"strings"
	"syscall"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *helloworld.String) (*helloworld.String, error) {
	if len(args.GetValue()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid parameters")
	}

	return &helloworld.String{Value: "Hello " + args.GetValue()}, nil
}

func (p *HelloServiceImpl) SayHello(*helloworld.String, helloworld.HelloService_SayHelloServer) error {
	return nil
}

func (p *HelloServiceImpl) Channel(stream helloworld.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		if err := stream.Send(&helloworld.String{Value: "Hello " + args.GetValue()}); err != nil {
			return err
		}
	}
}

func filter1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func filter2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}

func main() {
	if err := grpclb.Register("", "HelloService", "localhost", "8080", time.Second*10, 15); err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		s := <-ch
		log.Printf("recv signal: %q\n", s)
		grpclb.Unregister()
		os.Exit(1)
	}()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(filter1, filter2)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer()))
	helloworld.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "hello")
	})

	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			mux.ServeHTTP(w, r)
			return
		}

		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		}
	}))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.GracefulStop()
}
