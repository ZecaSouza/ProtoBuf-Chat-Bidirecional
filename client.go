package main

import (
	"bi-direcional/chat"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Conectar ao servidor gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor: %v", err)
	}
	defer conn.Close()

	// Criar cliente gRPC
	chatClient := chat.NewChatServiceClient(conn)

	// Criar contexto com cancelamento para encerrar a comunicação corretamente
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := chatClient.Join(ctx)
	if err != nil {
		log.Fatalf("Erro ao entrar no chat: %v", err)
	}
	defer stream.CloseSend()

	// Capturar nome do usuário
	fmt.Print("Digite seu nome: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user := scanner.Text()

	// Canal para capturar interrupções do sistema (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// Goroutine para receber mensagens do servidor
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Println("Conexão encerrada pelo servidor.")
				cancel()
				return
			}
			fmt.Printf("[%s] %s: %s\n", time.Unix(msg.Timestamp, 0).Format("15:04:05"), msg.User, msg.Text)
		}
	}()

	// Loop para enviar mensagens
	for {
		select {
		case <-sigChan:
			fmt.Println("\nEncerrando o chat...")
			cancel()
			return
		default:
			if scanner.Scan() {
				msg := &chat.Message{
					User:      user,
					Text:      scanner.Text(),
					Timestamp: time.Now().Unix(),
				}
				if err := stream.Send(msg); err != nil {
					log.Printf("Erro ao enviar mensagem: %v", err)
					return
				}
			} else {
				// Se scanner parar de ler (EOF), encerrar
				fmt.Println("\nSaindo do chat...")
				cancel()
				return
			}
		}
	}
}
