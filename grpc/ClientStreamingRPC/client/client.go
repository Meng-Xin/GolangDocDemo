package main

import (
	"CSRPC/protoc/pb"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"path/filepath"
	"strings"
	"time"
)

type Device interface {
	Report()
}

type device struct {
	Id     string
	Name   string
	Status string
}

func NewDevice(uuid string, name string) *device {
	return &device{
		Id:     uuid,
		Name:   name,
		Status: "正常",
	}
}

func (d *device) Update(stream pb.DeviceInformationCollectionService_SendStatusClient) {
	// 上报信息
	reportInfo := &pb.DeviceStatusUpdate{
		DeviceId:      d.Id,
		Name:          d.Name,
		StatusMessage: d.Status,
	}
	if err := stream.Send(reportInfo); err != nil {
		slog.Error("Error sending update", "Err:", err.Error())
		return
	}
	slog.Info("Device Report Info Success")
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDeviceInformationCollectionServiceClient(conn)
	stream, err := client.SendStatus(context.Background())
	if err != nil {
		return
	}
	// 假设这里的设备是Wifi
	wifiDevice := NewDevice(uuid.New().String(), "Wifi")

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	counter := 0
	for {
		select {
		case <-ticker.C:
			// 累计三次上报后模拟设备故障
			if counter == 3 {
				wifiDevice.Status = "故障"
			}
			// 例如这里发送的信息是设备自动上报状态，模拟每3秒发送一次更新
			wifiDevice.Update(stream)
			counter++
		}
	}
}

func getFileType(filePath string) string {
	// 使用 filepath 包获取文件的扩展名
	ext := filepath.Ext(filePath)

	// 去掉扩展名中的点号
	ext = strings.TrimPrefix(ext, ".")

	return ext
}
