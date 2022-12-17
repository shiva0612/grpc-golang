package main

import (
	"context"
	"fmt"
	"io"
	"log"
	um "shiva/proto/models/user"
	us "shiva/proto/services/user"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	client us.UserApiClient
)

func main() {

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client = us.NewUserApiClient(cc)

	unary()
	// cstream()
	// sstream()
	// bistream()

}

func unary() {
	ctx := context.Background()
	res, err := client.AddUser(ctx, &um.User{Id: "1", Email: []string{"shiva@example.com"}})
	if err != nil {
		log.Println("error in client-unary: ", err)
	}
	fmt.Println("unary response: ", res.Response)
}

func cstream() {
	ctx := context.Background()
	stream, err := client.AddUsers(ctx)
	if err != nil {
		log.Fatalln("error in client-stream: ", err)
	}

	for i := 0; i < 5; i++ {
		stream.Send(&um.User{
			Id: fmt.Sprintf("%v", i),
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while getting response from client stream: ", err)
	}

	fmt.Println("response from client stream: ", res.Response)
}

func sstream() {
	ctx := context.Background()
	stream, err := client.ListUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalln("error while getting response from server stream: ", err)
	}

	for {
		userres, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("error while receiving response from server stream: ", err)
		}
		fmt.Println(userres.Id)
	}
}

func bistream() {
	ctx := context.Background()
	resch := make(chan *um.User, 10)
	stream, err := client.ListTheseUsers(ctx)
	if err != nil {
		log.Println("error while requesting bi-directional stream: ", err)
	}
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending request for id = ", i)
			stream.Send(&um.StringRequest{Request: fmt.Sprintf("%v", i)})
		}
	}()
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("error while receiving response from bi stream: ", err)
		}
		fmt.Println("server response: ", user)
		resch <- user
	}
}
