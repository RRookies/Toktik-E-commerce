package grpcx

import (
	"Toktik-E-commerce/app/auth/utils/netx"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"strconv"
	"time"
)

type Server struct {
	*grpc.Server
	Port        int
	EtcdTTL     int64
	EtcdClient  *clientv3.Client
	etcdManager endpoints.Manager
	etcdKey     string
	cancel      func()
	Name        string
}

func (s *Server) Serve() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	port := strconv.Itoa(s.Port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return
	}
	err = s.register(ctx, port)
	if err != nil {
		return
	}
	slog.Info("用户服务开启成功，监听端口为:" + strconv.Itoa(s.Port))
	return s.Server.Serve(l)
}

func (s *Server) register(ctx context.Context, port string) (err error) {
	cli := s.EtcdClient
	serviceName := "service/" + s.Name
	em, err := endpoints.NewManager(cli, serviceName)
	if err != nil {
		return
	}
	s.etcdManager = em
	ip := netx.GetOutboundIP()

	s.etcdKey = serviceName + "/" + ip
	addr := ip + ":" + port
	leaseResp, err := cli.Grant(ctx, s.EtcdTTL)

	ch, err := cli.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return
	}
	go func() {
		for _ = range ch {

		}
	}()

	return em.AddEndpoint(ctx, s.etcdKey, endpoints.Endpoint{
		Addr: addr,
	}, clientv3.WithLease(leaseResp.ID))
}

func (s *Server) Close() (err error) {
	s.cancel()
	if s.etcdManager != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		err = s.etcdManager.DeleteEndpoint(ctx, s.etcdKey)
		if err != nil {
			return
		}
	}
	err = s.EtcdClient.Close()
	if err != nil {
		return
	}
	s.Server.GracefulStop()
	return
}
