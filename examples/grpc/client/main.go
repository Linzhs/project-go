package main

import (
	"context"
	"log"
	"project-go/pkg/grpclb"
	"project-go/pkg/rpc/pb/helloworld"
	"time"

	"google.golang.org/grpc/balancer/roundrobin"

	"google.golang.org/grpc/resolver"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
)

func main() {
	r := grpclb.NewResolver("", "HelloService")
	resolver.Register(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, r.Scheme()+":1234", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := helloworld.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &helloworld.String{Value: "grpc"})

	// handle grpc error
	st, ok := status.FromError(err)
	if !ok {
		log.Fatal(st.Err())
	}

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := stream.Send(&helloworld.String{Value: "hi"}); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}()

	log.Println(reply)
}
