package grpclb

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	etcd3 "github.com/coreos/etcd/clientv3"
)

var deregister = make(chan struct{})

// Register is service register
func Register(target, serviceName, host, port string, interval time.Duration, ttl int64) error {
	c, err := etcd3.New(etcd3.Config{Endpoints: strings.Split(target, ",")})
	if err != nil {
		return fmt.Errorf("grpclb: failed to new etcdv3 client: %w", err)
	}

	resp, err := c.Grant(context.TODO(), ttl)
	if err != nil {
		return fmt.Errorf("grpclb: failed to lease :%w", err)
	}

	var (
		serviceValue = net.JoinHostPort(host, port)
		serviceKey   = fmt.Sprintf("/%s/%s/%s", scheme, serviceName, serviceValue)
	)

	if _, err := c.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
		return fmt.Errorf("grpclb: set service %q with ttl to etcdv3 failed: %w", serviceName, err)
	}

	if _, err := c.KeepAlive(context.TODO(), resp.ID); err != nil {
		return fmt.Errorf("grpc:lb: refresh service %q with ttl to etcdv3 failed:%w", serviceName, err)
	}

	go func() {
		<-deregister
		c.Delete(context.Background(), serviceKey)
		deregister <- struct{}{}
	}()

	return nil
}

// UnRegister delete registered service from etcd
func Unregister() {
	deregister <- struct{}{}
	<-deregister
}
