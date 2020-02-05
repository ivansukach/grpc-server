package grpc_server

import (
	"github.com/ivansukach/grpc-server/protocol"
	"github.com/ivansukach/grpc-server/server"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	protocol.RegisterGetResponseServer(s, srv)

}
