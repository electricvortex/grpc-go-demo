package api

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "hi doe"}, nil
}
