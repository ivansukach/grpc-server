package main

import (
	"github.com/ivansukach/grpc-server/protocol"
	"github.com/ivansukach/grpc-server/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()                       //создали сервер
	srv := &server.GRPCServer{}                 //ссылка на пустую структуру
	protocol.RegisterGetResponseServer(s, srv)  //зарегистрировали сервер
	listener, err := net.Listen("tcp", ":1323") //просто слушаем порт
	s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
