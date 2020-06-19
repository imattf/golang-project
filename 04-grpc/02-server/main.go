package main

import (
	"context"
	"fmt"
	"net"

	"github.com/imattf/golang-project/04-grpc/02-server/echo"
	"google.golang.org/grpc"
)

//EchoServer is a ....
type EchoServer struct{}

//Echo is a...
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "My Echo: " + req.Message,
	}, nil
}

func main() {

	//create listener
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//create grps server
	s := grpc.NewServer()

	//create Echo server
	srv := &EchoServer{}

	//regsiter Echo server with the grpc service
	echo.RegisterEchoServerServer(s, srv)

	//tell grpc server to serve on our listener

	fmt.Println("Now serving at port 8080...")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}

}
