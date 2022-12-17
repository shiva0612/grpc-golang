package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	um "shiva/proto/models/user"
	us "shiva/proto/services/user"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserApi struct {
	us.UnimplementedUserApiServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	us.RegisterUserApiServer(s, &UserApi{})
	// bs.RegisterBookApiServer(s, &BookApi{})

	fmt.Println("server started ...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server stopeed: %v", err)
	}
}

// unary
func (*UserApi) AddUser(ctx context.Context, req *um.User) (*um.StringResponse, error) {
	fmt.Printf("in unary service: req = %v", req)

	res := &um.StringResponse{
		Response: "user with id = " + req.Id + " created",
	}
	return res, nil
}

// client streaming
func (*UserApi) AddUsers(stream us.UserApi_AddUsersServer) error {
	ids := make([]string, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := fmt.Sprintf("ids created: %v", ids)
			return stream.SendAndClose(&um.StringResponse{
				Response: res,
			})
		}
		if err != nil {
			log.Fatalln("error while listening to client streaming: %v", err)
		}
		fmt.Println("got request for: ", req.Id)
		ids = append(ids, req.Id)
	}
}

// server streaming
func (*UserApi) ListUsers(c *emptypb.Empty, stream us.UserApi_ListUsersServer) error {

	fmt.Println("request for listusers")
	for i := 0; i < 5; i++ {
		stream.Send(&um.User{
			Id: fmt.Sprintf("%v", i),
		})
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

// bi-streaming
func (*UserApi) ListTheseUsers(stream us.UserApi_ListTheseUsersServer) error {

	resch := make(chan um.User, 10)
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				close(resch)
				return
			}
			if err != nil {
				log.Fatalln("error in bidirectional stream: %v", err)
			}

			fmt.Println("received request for id = ", req.Request)
			resch <- um.User{Id: req.Request}
		}
	}()
	for res := range resch {
		stream.Send(&res)
	}
	return nil
}
