package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	hellopb "github.com/kazu697/grpc/src/pkg/grpc/src/api"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Name),
	}, nil
}

func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func NewServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// grpcサーバー起動
	s := grpc.NewServer()

	// grpcサーバーにGreetingServiceを登録する
	hellopb.RegisterGreetingServiceServer(s, NewServer())

	// grpcurlを実行できるようにするためにreflectionを登録
	reflection.Register(s)

	// 作成したgrpcサーバーを8080番ポートで起動
	go func() {
		log.Printf("start GRPC Server port %d", port)
		s.Serve(listener)
	}()

	// Ctrl + Cで終了
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping GRPC Server...")
	s.GracefulStop()
}
