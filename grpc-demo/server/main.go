package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"grpc-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	// Имитация долгой операции, чтобы проверить GracefulStop
	time.Sleep(2 * time.Second)
	return &proto.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	// Канал для ожидания сигнала завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Server started on :50051")

	// Запускаем сервер в отдельной горутине
	go func() {
		// Serve — блокирующий вызов, поэтому он в горутине.
		// Если вы используете версию gRPC, где есть ServeAsync, можно использовать его.
		if err := s.Serve(lis); err != nil {
			log.Printf("server failed to serve: %v", err)
		}
	}()

	// Ожидаем сигнал прерывания
	<-stop
	log.Println("Shutdown signal received")

	// Останавливаем сервер.
	// GracefulStop дождется завершения активных запросов (в нашем примере — те, что спят 2 секунды).
	// На новые запросы сервер отвечать уже не будет.
	s.GracefulStop()

	// Если у вас есть пул соединений (*sql.DB), горутины или другие ресурсы,
	// их контекст нужно отменить здесь, а ресурсы — закрыть.

	log.Println("Server gracefully stopped")
}
