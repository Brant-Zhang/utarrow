package diskqueue

import (
	"testing"
)

const sample = "today is good,tomorrow will be better4"

func TestWrite(t *testing.T) {
	db := New("./", "hangzhou")
	defer db.Close()
	err := db.Put([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
}
