package main

import (
	"bi-direcional/chat"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChatServer struct {
	chat.UnimplementedChatServiceServer
	mu       sync.Mutex
	clients  map[chat.ChatService_JoinServer]bool
	messages chan *chat.Message
}

func NewChatServer() *ChatServer {
	server := &ChatServer{
		clients:  make(map[chat.ChatService_JoinServer]bool),
		messages: make(chan *chat.Message, 100), // Buffer para evitar bloqueios
	}

	go server.broadcastMessages()
	return server
}

// Goroutine para enviar mensagens a todos os clientes
func (s *ChatServer) broadcastMessages() {
	for msg := range s.messages {
		s.mu.Lock()
		for client := range s.clients {
			if err := client.Send(msg); err != nil {
				log.Printf("Erro ao enviar mensagem: %v", err)
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}

// Método Join (stream bidirecional)
func (s *ChatServer) Join(stream chat.ChatService_JoinServer) error {
	// Registrar cliente
	s.mu.Lock()
	s.clients[stream] = true
	s.mu.Unlock()

	defer func() {
		// Remover cliente ao desconectar
		s.mu.Lock()
		delete(s.clients, stream)
		s.mu.Unlock()
	}()

	// Loop para receber mensagens do cliente
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("Cliente desconectado: %v", err)
			return err // Encerra o stream corretamente
		}

		// Encaminhar mensagem para broadcast
		s.messages <- msg
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar servidor: %v", err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, NewChatServer())

	log.Println("Servidor gRPC rodando na porta 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Erro ao rodar servidor: %v", err)
	}
}