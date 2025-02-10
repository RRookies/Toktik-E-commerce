package grpcs

import (
	"context"
	"net"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
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
		return err
	}
	err = s.register(ctx, port)
	if err != nil {
		return err
	}
	return s.Server.Serve(l)
}

func (s *Server) register(ctx context.Context, port string) (err error) {
	client := s.EtcdClient
	serviceName := "service/" + s.Name
	em, err := endpoints.NewManager(client, serviceName)
	if err != nil {
		return err
	}
	s.etcdManager = em
	ip := "localhost"

	s.etcdKey = serviceName + "/" + ip
	addr := ip + ":" + port
	leaseResp, err := client.Grant(ctx, s.EtcdTTL)
	if err != nil {
		return
	}
	ch, err := client.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return
	}
	go func(ctx context.Context) {
		for {
			select {
			case <-ch:

			case <-ctx.Done():
				return
			}
		}
	}(ctx)

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
