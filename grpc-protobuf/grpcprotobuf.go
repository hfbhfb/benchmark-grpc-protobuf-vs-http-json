package grpcprotobuf

import (
	"errors"
	"log"
	"net"
	"net/mail"
	"time"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	FlagSleep = false
)

// Start entrypoint
func StartDelay() {
	lis, _ := net.Listen("tcp", ":80")

	srv := grpc.NewServer()
	proto.RegisterAPIServer(srv, &Server{})
	log.Println(srv.Serve(lis))
}

// Start entrypoint
func Start() {
	lis, _ := net.Listen("tcp", ":60000")

	srv := grpc.NewServer()
	proto.RegisterAPIServer(srv, &Server{})
	log.Println(srv.Serve(lis))
}

// Server type
type Server struct{}

// CreateUser handler
func (s *Server) CreateUser(ctx context.Context, in *proto.User) (*proto.Response, error) {
	validationErr := validate(in)
	if validationErr != nil {
		return &proto.Response{
			Code:    500,
			Message: validationErr.Error(),
		}, validationErr
	}

	in.Id = "1000000"
	if FlagSleep {
		time.Sleep(10 * time.Millisecond)
	}

	return &proto.Response{
		Code:    200,
		Message: "OK",
		User:    in,
	}, nil
}

func validate(in *proto.User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
