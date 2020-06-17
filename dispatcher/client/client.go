package client

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/branthz/etcd/clientv3"
	"github.com/branthz/utarrow/lib/log"
)

func init() {
	log.Setup("", "debug")
}

type Meta struct {
	IP     string
	Load   uint32 //系统目前的负载数量,用在平衡用户端申请资源时平衡cluster内各节点分配
	Weight uint8  //每个节点的权值，以1为单位；1的倍数;后期可以改成小数.
	//ID     uint64 //leaseid as a unique id for etcd
}

type node struct {
	data    Meta
	leaseid clientv3.LeaseID //节点租期
	stop    chan error       //节点主动退出
	conn    *clientv3.Client
	service string
}

func NewNode(weight uint8, name string, ip string, etcd []string) *node {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd,
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	m := &Meta{
		IP:     ip,
		Load:   0,
		Weight: weight,
	}
	//m.ID = uint64(time.Now().Unix())
	n := &node{
		data:    *m,
		service: name,
		stop:    make(chan error),
		conn:    cli,
	}
	return n
}

func (s *node) Stop() {
	s.stop <- nil
}

func (s *node) gettest(key string) error {
	res, err := s.conn.Get(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Info("get response:%v\n", res.Kvs)
	return nil
}

func (s *node) revoke() error {
	_, err := s.conn.Revoke(context.TODO(), s.leaseid)
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("servide:%s stop\n", s.service)
	return err
}

func (s *node) Start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	for {
		select {
		case err := <-s.stop:
			s.revoke()
			return err
		case <-s.conn.Ctx().Done():
			return errors.New("node node")
		case ka, ok := <-ch:
			if !ok {
				log.Infoln("keep alive channel closed")
				//TODO
				//reconnect to etcd node
				s.revoke()
				return nil
			} else {
				log.Info("Recv reply from service: %s, ttl:%d", s.service, ka.TTL)
			}
		}
	}
}

func (s *node) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resp, err := s.conn.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatalln(err)
	}
	value, _ := json.Marshal(s.data)
	s.conn.Put(context.TODO(), s.service, string(value), clientv3.WithLease(resp.ID))
	s.leaseid = resp.ID
	return s.conn.KeepAlive(context.TODO(), resp.ID)
}
