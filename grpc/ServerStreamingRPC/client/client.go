package main

import (
	"SSRPC/protoc/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8009", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewStockServiceClient(conn)
	stream, err := client.GetStockById(context.Background(), &pb.StockReq{StockId: 100001})
	if err != nil {
		return
	}

	// Steaming 流是源源不断的，因此需要持续接受。
	for {
		resp, err := stream.Recv()
		if err != nil {
			slog.Error("failed client accept msg", "Err:", err.Error())
			break
		}
		slog.Info("Client Accept ServerStreaming……", "Msg:", resp.GetData())
	}

}
