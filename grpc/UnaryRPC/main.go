package main

import (
	"UnaryRPC/protoc/pb"
	context "context"
	"errors"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sync"
	"time"
)

type data struct {
	MinHeat   float32 // 温度
	MaxHeat   float32 // 温度
	Status    uint32  // 天气状态
	WindSpeed uint32  // 风速
	Time      int64   // 时间戳
	City      string  // 城市
}

type Weather struct {
	Weathers map[string]data
	mutex    sync.Mutex
}

func (w *Weather) GetWeatherData(city string) (data, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if val, ok := w.Weathers[city]; ok {
		return val, nil
	}
	return data{}, errors.New("未找到该城市")
}

var GlobalWeather = new(Weather)

// WeatherImpl 该类用于实现WeatherServiceServer接口
type WeatherImpl struct {
	// 必须嵌入 UnimplementedWeatherServiceServer ,该接口是对应方法的未实现调用。
	pb.UnimplementedWeatherServiceServer
}

// GetTodayData 获取天气数据
func (w WeatherImpl) GetTodayData(ctx context.Context, req *pb.WeatherReq) (*pb.WeatherResp, error) {
	// 查询城市天气
	weatherData, err := GlobalWeather.GetWeatherData(req.City)
	if err != nil {
		return nil, err
	}
	// 构建返回数据
	resp := &pb.WeatherResp{Data: &pb.Weather{
		MinHeat:   weatherData.MinHeat,
		MaxHeat:   weatherData.MaxHeat,
		Status:    weatherData.Status,
		WindSpeed: weatherData.WindSpeed,
		Time:      weatherData.Time,
		City:      weatherData.City,
	}}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:7009")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer listen.Close()
	// 实例化天气信息
	GlobalWeather.Weathers = map[string]data{
		"bj": {12.8, 22, 1, 7, time.Now().Unix(), "北京"},
		"sh": {20, 29, 2, 2, time.Now().Unix(), "上海"},
		"hz": {18.6, 27, 3, 5, time.Now().Unix(), "杭州"},
		"sz": {21, 30, 4, 3, time.Now().Unix(), "深圳"},
	}
	// 实例化GRPC
	grpcServer := grpc.NewServer()
	// 注册服务
	pb.RegisterWeatherServiceServer(grpcServer, &WeatherImpl{})
	// 绑定Listen到GRPC
	if err = grpcServer.Serve(listen); err != nil {
		slog.Error("failed to grpcServer", "Err:", err.Error())
	}
}
