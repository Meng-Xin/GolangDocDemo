package main

import (
	"UnaryRPC_Interceptor/protoc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7009", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer conn.Close()
	client := pb.NewWeatherServiceClient(conn)
	bjData, err := client.GetTodayData(context.Background(), &pb.WeatherReq{City: "bj"})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(bjData.GetData())

	shData, err := client.GetTodayData(context.Background(), &pb.WeatherReq{City: "sh"})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(shData.GetData())

	hzData, err := client.GetTodayData(context.Background(), &pb.WeatherReq{City: "hz"})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(hzData.GetData())

	szData, err := client.GetTodayData(context.Background(), &pb.WeatherReq{City: "sz"})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(szData.GetData())
}
