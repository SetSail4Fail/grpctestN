package main

import (
	"context"
	test "grpctestN/grpcGenerated" // Замените на путь к вашему сгенерированному коду
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Реализация сервера
type server struct {
	test.UnimplementedGreeterServer // Встроенная структура для совместимости
}

// Реализация метода SayHello
func (s *server) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &test.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	// Создаем TCP-листенер на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем gRPC-сервер
	s := grpc.NewServer()

	// Регистрируем наш сервис на сервере
	test.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	// Запускаем сервер
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
