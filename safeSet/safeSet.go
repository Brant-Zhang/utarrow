//with time expire triger for entry
//get entry slice alive
//multi thread safe
package safeset

import (
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type valueProto interface{}
type entryProto struct {
	v     valueProto
	utime int64
}

type mapProto map[string]entryProto
type setProto []entryProto

type set_st struct {
	m               mapProto
	mu              sync.Mutex
	set             setProto
	expireTime      int //unit is second
	triggerInterval int
	flag            int //for cancel time expirtion
	pset            unsafe.Pointer
}

//support expire entry and no expire
func Newset(flag, ex, tr int) *set_st {
	s := new(set_st)
	s.m = make(mapProto)
	s.set = make(setProto, 0)
	s.pset = unsafe.Pointer(&s.set)

	s.expireTime = ex
	s.triggerInterval = tr
	s.flag = flag
	go s.refresh()
	return s
}

func (this *set_st) Add(k string, v valueProto) {
	var entry entryProto
	entry.v = v
	entry.utime = time.Now().Unix()
	this.mu.Lock()
	this.m[k] = entry
	this.mu.Unlock()
}

func (this *set_st) Delete(k string) {
	this.mu.Lock()
	delete(this.m, k)
	this.mu.Unlock()
}

func (this *set_st) refresh() {
	t1 := time.NewTicker(time.Second * time.Duration(this.triggerInterval))
	for {
		<-t1.C
		tm := time.Now().Unix()
		set := make(setProto, 0)
		this.mu.Lock()
		for k, v := range this.m {
			if tm-v.utime > int64(this.expireTime) {
				delete(this.m, k)
			} else {
				set = append(set, v)
			}
		}
		this.mu.Unlock()
		this.updateSet(set)
	}
}

func (this *set_st) GetSet() setProto {
	return *(*setProto)(atomic.LoadPointer(&this.pset))
}

func (this *set_st) updateSet(s setProto) {
	atomic.StorePointer(&this.pset, unsafe.Pointer(&s))
}
