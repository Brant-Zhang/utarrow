package dispatcher

import (
	"log"
	"time"

	"github.com/branthz/etcd/clientv3"
)

//现在纠结主从节点的切换问题，如何标记主从节点（eg通过节点启动时分配信息？先后顺序？）
//假设按照先后顺序，先连接etcd-room的打上master标签，后连接的打上slave标签
//某个时间点master停了，etcd监听到该事件，将slave提升为master
//上述信息需要etcd管理不方便
//服务节点连接etcd先获取room下的集群信息，没有将自身打上master标签；已经有了将自身打上slave标签；
//每个服务监听room节点数，当发生变化时查看是否还有master；如没有将自身提升为master

type backend struct {
	load      int64
	partition string
	master    string
	slave     string
}

type backends struct {
	list []backend
}

func (b *backends) save(partition string, role string, ip string) {
	for i := 0; i < len(b.list); i++ {
		if b.list[i].partition == partition {
			if role == "master" {
				//b.list[i].master
			}
		}
	}
}

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
	//TODO
	//针对该key是否已经配置过了
	//是返回
	//否，获取服务列表，负载最低优先
	return ""
}

//TODO
//func (p *pacher) GetServerList() {
//	res, err := p.Conn.Get(context.TODO(), "/controller/part", clientv3.WithPrefix())
//	if err != nil {
//		log.Fatalln(err)
//		return
//	}
//	ls := new(backends)
//	for v := range res.Kvs {
//		ss := strings.Split(v.Key, "/")
//
//		v.Value
//	}
//}

func (p *pacher) watchNode() {
	//rch := p.Conn.Watch(context.Background(), p.Path, clientv3.WithPrefix())
}
