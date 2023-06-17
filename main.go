package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"

	"dice/pb"

	"google.golang.org/grpc"
)

// わざわざ myServer を定義せずに、 pb.UnimplementedDiceServiceServer を使えば良さそうに見えるが、
// `func (s *myServer) RollDice` 部分で `cannot define new methods on non-local type` エラーが発生してしまう
// (外部ライブラリの型に対して、メソッドを定義することはできないよ！エラー)
// そのため、以下のような冗長な書き方をあえてする必要がある
type myServer struct {
	pb.UnimplementedDiceServiceServer
}

func getRandomInt(a int, b int) int {
	return rand.Intn(b-a+1) + a
}

func (s *myServer) RollDice(ctx context.Context, req *pb.RollDiceRequest) (*pb.RollDiceResponse, error) {
	randomNum := getRandomInt(
		int(req.MinNum),
		int(req.MaxNum),
	)

	return &pb.RollDiceResponse{
		Result: int32(randomNum),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterDiceServiceServer(server, &myServer{})

	go func() {
		log.Println("start gRPC server port: 50051")
		server.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
