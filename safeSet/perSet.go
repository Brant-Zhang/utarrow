package safeset

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type Set_pst struct {
	m    mapProto
	mu   sync.RWMutex
	set  setProto
	pset unsafe.Pointer
}

func Newpset() *Set_pst {
	s := new(Set_pst)
	s.m = make(mapProto)
	s.set = make(setProto, 0)
	s.pset = unsafe.Pointer(&s.set)
	return s
}

func (this *Set_pst) Add(k string, v valueProto) {
	var entry entryProto
	entry.v = v
	set := make(setProto, 0)

	this.mu.Lock()
	this.m[k] = entry
	this.mu.Unlock()
	this.mu.RLock()
	for _, v := range this.m {
		set = append(set, v)
	}
	this.mu.RUnlock()
	this.updateSet(set)
}

func (this *Set_pst) Delete(k string) {
	this.mu.Lock()
	delete(this.m, k)
	this.mu.Unlock()

	set := make(setProto, 0)
	this.mu.RLock()
	for _, v := range this.m {
		set = append(set, v)
	}
	this.mu.RUnlock()
	this.updateSet(set)
}

func (this *Set_pst) GetSet() setProto {
	return *(*setProto)(atomic.LoadPointer(&this.pset))
}

func (this *Set_pst) updateSet(s setProto) {
	atomic.StorePointer(&this.pset, unsafe.Pointer(&s))
}
