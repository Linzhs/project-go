package grpclb

import (
	"context"
	"fmt"
	"strings"

	mvccpb2 "github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/mvcc/mvccpb"

	"google.golang.org/grpc/resolver"

	etcd3 "github.com/coreos/etcd/clientv3"
)

const (
	scheme = "etcdv3_resolver"
)

// resolver is the implementation of grpc.naming.Resolver
type Resolver struct {
	target      string
	serviceName string
	c           *etcd3.Client
	cc          resolver.ClientConn
}

// NewResolver return resolver builder
// target example: "http://127.0.0.1:2379,http://127.0.0.1:12379,http://127.0.0.1:22379"
func NewResolver(target, serviceName string) resolver.Builder {
	return &Resolver{serviceName: serviceName, target: target}
}

func (r *Resolver) Scheme() string {
	return scheme
}

func (r *Resolver) ResolveNow(rn resolver.ResolveNowOption) {

}

func (r *Resolver) Close() {

}

// Build to resolver.Resolver
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (_ resolver.Resolver, err error) {
	r.c, err = etcd3.New(etcd3.Config{Endpoints: strings.Split(r.target, ",")})
	if err != nil {
		err = fmt.Errorf("grpclb: failed to create etcdv3 client: %w", err)
		return
	}

	r.cc = cc
	go r.watch(fmt.Sprintf("/%s/%s/", r.Scheme(), r.serviceName))

	return r, nil
}

// watch service discovery
func (r *Resolver) watch(prefix string) {
	addrDict := make(map[string]resolver.Address)

	fnUpdate := func() {
		addrs := make([]resolver.Address, 0, len(addrDict))
		for _, v := range addrDict {
			addrs = append(addrs, v)
		}

		r.cc.NewAddress(addrs)
	}

	resp, err := r.c.Get(context.Background(), prefix, etcd3.WithPrefix())
	if err == nil {
		for _, v := range resp.Kvs {
			addrDict[string(v.Value)] = resolver.Address{Addr: string(v.Value)}
		}
	}

	fnUpdate()

	ch := r.c.Watch(context.Background(), prefix, etcd3.WithPrefix(), etcd3.WithPrevKV())
	for v := range ch {
		for _, ev := range v.Events {
			switch ev.Type {
			case mvccpb2.Event_EventType(mvccpb.PUT):
				addrDict[string(ev.Kv.Key)] = resolver.Address{Addr: string(ev.Kv.Value)}
			case mvccpb2.Event_EventType(mvccpb.DELETE):
				delete(addrDict, string(ev.PrevKv.Key))
			}
		}

		fnUpdate()
	}
}
