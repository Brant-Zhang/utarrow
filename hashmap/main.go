//golang hashmap的实现
package main

import (
	"fmt"
	"unsafe"
)

const (
	// Maximum number of key/value pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	// Maximum average load of a bucket that triggers growth.
	loadFactor = 6.5

	// Maximum key or value size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	// Fast versions cannot handle big values - the cutoff size for
	// fast versions in ../../cmd/internal/gc/walk.go must be at most this value.
	maxKeySize   = 128
	maxValueSize = 128

	// data offset should be the size of the bmap struct, but needs to be
	// aligned correctly.  For amd64p32 this means 64-bit alignment
	// even though pointers are 32 bit.
	dataOffset = unsafe.Offsetof(struct {
		b bmap
		v int64
	}{}.v)

	// Possible tophash values.  We reserve a few possibilities for special marks.
	// Each bucket (including its overflow buckets, if any) will have either all or none of its
	// entries in the evacuated* states (except during the evacuate() method, which only happens
	// during map writes and thus no one else can observe the map during that time).
	empty          = 0 // cell is empty
	evacuatedEmpty = 1 // cell is empty, bucket is evacuated.
	evacuatedX     = 2 // key/value is valid.  Entry has been evacuated to first half of larger table.
	evacuatedY     = 3 // same as above, but evacuated to second half of larger table.
	minTopHash     = 4 // minimum tophash for a normal filled cell.

	// flags
	iterator    = 1 // there may be an iterator using buckets
	oldIterator = 2 // there may be an iterator using oldbuckets
	hashWriting = 4 // a goroutine is writing to the map

	// sentinel bucket ID for iterator checks
	noCheck = 1<<(8*8) - 1
)

type hmap struct {
	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
	// ../reflect/type.go.  Don't change this structure without also changing that code!
	count int // # live cells == size of map.  Must be first (used by len() builtin)
	flags uint8
	B     uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	hash0 uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.overflow.
	// Overflow is used only if key and value do not contain pointers.
	// overflow[0] contains overflow buckets for hmap.buckets.
	// overflow[1] contains overflow buckets for hmap.oldbuckets.
	// The first indirection allows us to reduce static size of hmap.
	// The second indirection allows to store a pointer to the slice in hiter.
	overflow *[2]*[]*bmap
}

// A bucket for a Go map.
type bmap struct {
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

func newarray() {
	p := make([]byte)
}

func makemap(hint int) *hmap {
	// find size parameter which will hold the requested # of elements
	// cal the fit map size
	B := uint8(0)
	for ; hint > bucketCnt && float32(hint) > loadFactor*float32(uintptr(1)<<B); B++ {
	}

	buckets = newarray(t.bucket, uintptr(1)<<B)
}

func main() {

}
