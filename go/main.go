package main

import (
	"context"
	"fmt"
	"net/http"

	hello_worldv1 "contain-yourself/gen/go/hello_world/v1"
	"contain-yourself/gen/go/hello_world/v1/hello_worldv1connect"

	"connectrpc.com/connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := hello_worldv1connect.NewHelloWorldServiceHandler(&helloWorldServer{})
	mux.Handle(path, handler)
	fmt.Println("Server is running on", address)
	http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// helloWorldServer implements the HelloWorld service.
type helloWorldServer struct {
	hello_worldv1connect.UnimplementedHelloWorldServiceHandler
}

// SayHello sends a greeting.
func (s *helloWorldServer) SayHello(
	ctx context.Context,
	req *connect.Request[hello_worldv1.HelloRequest],
) (*connect.Response[hello_worldv1.HelloResponse], error) {
	name := req.Msg.Name
	if name == "" {
		name = "World"
	}
	response := &hello_worldv1.HelloResponse{
		Message: "Hello, " + name + "!",
	}
	return connect.NewResponse(response), nil
}
