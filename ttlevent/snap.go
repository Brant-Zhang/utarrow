package ttlevent

import (
	"encoding/gob"
	"os"
)

const dir = "/home/app/meta/"
const filename = "ttl.data"

func (c *Cache) backup() error {
	f, err := os.OpenFile(dir+filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	enc.Encode(c.pool)
	return nil
}

func (c *Cache) restore() {
	fr, err := os.OpenFile(dir+filename, os.O_RDONLY, 0755)
	if err != nil {
		return
	}
	dec := gob.NewDecoder(fr)
	dec.Decode(&c.pool)
	fr.Close()
	return
}
