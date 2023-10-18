package main

import (
	"SSRPC/protoc/pb"
	"errors"
	"google.golang.org/grpc"
	"log/slog"
	"math/rand"
	"net"
	"sync"
	"time"
)

type stock struct {
	Name             string
	TradingVolume    int64
	Price            float32
	IncrOrDecrAmount float32
	IncrOrDecrRate   float32
	Timestamp        int64
	mutex            *sync.RWMutex
}

func (s *stock) computeStock(amount float32, timing time.Time) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 计算涨跌额度
	s.IncrOrDecrAmount = s.Price - amount
	// 计算涨跌率
	s.IncrOrDecrRate = (s.Price - amount) / s.Price
	// 计算最新价格
	s.Price = amount
	// 成交率
	s.TradingVolume++
	// 时间戳
	s.Timestamp = timing.Unix()
}

type GlobalStock struct {
	stocks map[int64]*stock
	mutex  sync.Mutex
}

func (g *GlobalStock) getStockById(id int64) (*stock, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if v, ok := g.stocks[id]; ok {
		return v, nil
	}
	return nil, errors.New("未找到对应的股票信息")
}

var StockManager = new(GlobalStock)

type StockServiceImpl struct {
	pb.UnimplementedStockServiceServer
}

func (s StockServiceImpl) GetStockById(req *pb.StockReq, stream pb.StockService_GetStockByIdServer) error {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			data, err := StockManager.getStockById(req.StockId)
			if err != nil {
				return err
			}
			// 组装信息
			msg := &pb.StockResp{Data: &pb.Stock{
				Name:             data.Name,
				Price:            data.Price,
				TradingVolume:    data.TradingVolume,
				IncrOrDecrAmount: data.IncrOrDecrAmount,
				IncrOrDecrRate:   data.IncrOrDecrRate,
				Timestamp:        data.Timestamp,
			}}
			// 发送Stream流
			if err = stream.SendMsg(msg); err != nil {
				slog.Error("failed serverStreamingRPC sendMsg", "Err:", err.Error())
				return err
			}
		}
	}
}

func main() {
	// 初始化股票信息
	StockManager.stocks = map[int64]*stock{
		100001: &stock{"网易", 2000000, 3.32, 0.0, 0, time.Now().Unix(), &sync.RWMutex{}},
		100002: &stock{"腾讯", 4000000, 4.32, 0.0, 0, time.Now().Unix(), &sync.RWMutex{}},
		100003: &stock{"完美世界", 6000000, 2.32, 0.0, 0, time.Now().Unix(), &sync.RWMutex{}},
		100004: &stock{"米哈游", 8000000, 3.78, 0.0, 0, time.Now().Unix(), &sync.RWMutex{}},
	}

	// 更新股票数据
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case timing := <-tick.C:
				StockManager.mutex.Lock()
				for _, v := range StockManager.stocks {
					// 随机主题 0 为正值 1 为负值
					randTeam := rand.Intn(2)
					// 随机amount 股票便宜值
					randAmount := rand.Float32()
					if randTeam == 1 {
						randAmount = -randAmount
					}
					// 计算股票最新数据
					v.computeStock(randAmount, timing)
				}
				StockManager.mutex.Unlock()
			}
		}
	}()
	// 实现GRPC
	listen, err := net.Listen("tcp", "127.0.0.1:8009")
	if err != nil {
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStockServiceServer(grpcServer, &StockServiceImpl{})

	if err = grpcServer.Serve(listen); err != nil {
		slog.Error("failed grpc server ", "Err:", err.Error())
	}
}
