package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server is a server model
type Server struct {

}

// SayHello is service mettoh
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received %s",message.Body)
	return &Message{Body: "Hello "+message.Body+"!"}, nil
}