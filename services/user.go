package services

import (
	"context"
	"fmt"
	"time"

	"github.com/brunocicom/fc_grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	newUser := pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User:   &newUser,
	})

	// changing id
	newUser.Id = "124"

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User:   &newUser,
	})

	time.Sleep(time.Second * 3)

	return nil
}
