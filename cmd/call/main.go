package main

import (
	"context"
	"fmt"

	"github.com/dsrvlabs/vatz-proto/manager/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Call Plugin")

	cred := insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}

	cli := v2.NewRequestHandlerClient(conn)

	ctx := context.Background()
	req := &v2.UserRequest{
		Plugin: "snippet.grpc.reflection.HelloService",
		Method: "Hello",
		Fields: []*v2.FieldSpec{
			{
				Name: "name",
				Type: "string",
				Value: "rootwarp",
			},
			{
				Name: "age",
				Type: "int32",
				Value: "30",
			},
		},
	}

	resp, err := cli.SendRequest(ctx, req)

	fmt.Printf("%+v, %+v\n", resp, err)
}
