package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceLoader interface {
	ServiceRegistry() error
	deregisterService() error
}

type EtcdServiceRegistry struct {
	client     *clientv3.Client
	addr       string
	serverName string
	instanceId string
	grantResp  *clientv3.LeaseGrantResponse
}

func NewEtcdServiceRegistry(client *clientv3.Client, etcdaddr string, serverName string) (*EtcdServiceRegistry, error) {
	return &EtcdServiceRegistry{
		client:     client,
		addr:       etcdaddr,
		serverName: serverName,
		instanceId: uuid.New().String(),
		grantResp:  nil,
	}, nil
}
func (esr *EtcdServiceRegistry) ServiceRegistry() error {
	// 创建租约
	grantResp, err := esr.client.Grant(context.Background(), 100)
	if err != nil {
		return errors.New("创建租约失败")
	}
	key := fmt.Sprintf("/service/%s/%s", esr.serverName, esr.instanceId)
	_, err = esr.client.Put(context.Background(), key, esr.addr, clientv3.WithLease(grantResp.ID))

	if err != nil {
		return errors.New("创建键值失败")
	}
	esr.client.KeepAlive(context.Background(), grantResp.ID)
	esr.grantResp = grantResp
	return nil
}

func (esr *EtcdServiceRegistry) DeregisterService() error {
	key := fmt.Sprintf("/service/%s/%s", esr.serverName, esr.instanceId)
	_, err := esr.client.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}
