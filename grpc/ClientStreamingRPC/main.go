package main

import (
	"CSRPC/protoc/pb"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

type DeviceInfoCollectionServiceImpl struct {
	pb.UnimplementedDeviceInformationCollectionServiceServer
}

func (d DeviceInfoCollectionServiceImpl) SendStatus(stream pb.DeviceInformationCollectionService_SendStatusServer) error {
	for {
		report, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving status update: %v", err)
			return err
		}
		log.Printf("Received status from device %s: %s：%s", report.DeviceId, report.Name, report.StatusMessage)
		// 在实际应用中，可以在这里处理设备状态信息， 例如这里判断设备故障进行报警。
		if report.StatusMessage == "故障" {
			slog.Error("Warning: device is fault", "DeviceId:", report.DeviceId, "DeviceName:", report.Name, "DeviceStatus:", report.StatusMessage)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDeviceInformationCollectionServiceServer(s, &DeviceInfoCollectionServiceImpl{})
	log.Println("Server is listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
