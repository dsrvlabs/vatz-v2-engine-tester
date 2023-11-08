package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dsrvlabs/vatz-v2-engine-tester/proto/hello"
	pb "github.com/dsrvlabs/vatz-proto/manager/v2"
)

type helloServer struct {
	hello.UnimplementedHelloServiceServer
}

func (s *helloServer) Hello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("Hello", in.Name, in.Age)

	return &hello.HelloResponse{
		GreetingMsg: fmt.Sprintf("Hey %s(%d)", in.Name, in.Age),
	}, nil
}

func main() {
	fmt.Println("Start plugin server")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go startPluginServer(&wg)
	go startRegistration(&wg)

	wg.Wait()
}

func startPluginServer(wg *sync.WaitGroup) {
	fmt.Println("startPluginServer")

	defer wg.Done()

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &helloServer{})
	reflection.Register(s)

	l, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(l); err != nil {
		panic(err)
	}
}

func startRegistration(wg *sync.WaitGroup) {
	fmt.Println("startRegistration")

	defer wg.Done()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cli := pb.NewRegistrationServiceClient(conn)

	ctx := context.Background()
	resp, err := cli.RegisterPlugin(ctx, &pb.RegisterRequest{
		Name:    "snippet.grpc.reflection.HelloService", // TODO:
		Address: "127.0.0.1",
		Port:    9090,
	})

	fmt.Println(resp, err)
}
