//1. node-----etcd keepalive
//2. controller -----etcd keepalive
//3. 重连
//4. fencing token

package nfence

import (
	"context"
	"encoding/json"
	"time"

	"github.com/branthz/etcd/clientv3"
)

const (
	PathPrefix = "nfence/"
)

type ninfo struct {
}

type edata interface {
	encode() []byte
}

type Zclient struct {
	ID      string
	client  *clientv3.Client
	Info    edata
	leaseid clientv3.LeaseID
	stop    chan error
}

type memInfo struct {
	host string
	id   string
	tm   int64
}

func (m memInfo) encode() []byte {
	d, _ := json.Marshal(m)
	return d
}

func NewZclient(etcdServer []string, id string, host string) *Zclient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdServer,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return &Zclient{
		ID:     id,
		client: cli,
		Info:   memInfo{host: host, id: id, tm: time.Now().Unix()},
		stop:   make(chan error),
	}
}

func (s *Zclient) run() {
	ch, err := s.keepAlive()
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-s.stop:
			s.revoke()
		case <-s.client.Ctx().Done():
			Error("server closed")
			return
		case ka, ok := <-ch:
			if !ok {
				Debug("keep alive channel closed")
				s.revoke()
			} else {
				Debugf("Recv reply from service: %s, ttl:%d", s.ID, ka.TTL)
			}
		}
	}
}

func (s *Zclient) Stop() {
	s.stop <- nil
}

func (s *Zclient) revoke() error {
	_, err := s.client.Revoke(context.TODO(), s.leaseid)
	if err != nil {
		Fatal(err)
	}
	Infof("servide:%s stop\n", s.ID)
	return err
}

func (s *Zclient) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resp, err := s.client.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	key := PathPrefix + s.ID
	value := s.Info.encode()
	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
	if err != nil {
		return nil, err
	}
	s.leaseid = resp.ID
	return s.client.KeepAlive(context.TODO(), resp.ID)
}
