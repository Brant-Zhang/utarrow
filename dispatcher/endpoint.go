package dispatcher

import (
	"log"
	"time"

	"github.com/branthz/etcd/clientv3"
)

type pacher struct {
	Path string
	Conn *clientv3.Client
}

func NewPacher(etcd string, watchPath string) (*pacher, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcd},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	p := &pacher{
		Path: watchPath,
		Conn: cli,
	}
	go p.watchNode()
	return p, nil
}

//根据终端分配一个server地址
//如果之前不存在，根据集群里已经存在数目返回利用率最小的
//如果之前已经存在,则回去存在的值
func (p *pacher) Get(endpoint string) string {
	return ""
}

func (p *pacher) watchNode() {
	//rch := p.Conn.Watch(context.Background(), p.Path, clientv3.WithPrefix())
}
