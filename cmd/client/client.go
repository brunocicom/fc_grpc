package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/brunocicom/fc_grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Bruno",
		Email: "bruno@email.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Bruno",
		Email: "bruno@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := getMockUsers()

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func getMockUsers() []*pb.User {
	return []*pb.User{
		&pb.User{
			Id:    "b1",
			Name:  "bruno",
			Email: "bruno@email.com",
		},
		&pb.User{
			Id:    "b2",
			Name:  "bruno 2",
			Email: "bruno2@email.com",
		},
		&pb.User{
			Id:    "b3",
			Name:  "bruno 3",
			Email: "bruno3@email.com",
		},
		&pb.User{
			Id:    "b4",
			Name:  "bruno 4",
			Email: "bruno4@email.com",
		},
		&pb.User{
			Id:    "v5",
			Name:  "bruno 5",
			Email: "bruno5@email.com",
		},
	}
}
