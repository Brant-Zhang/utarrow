package diskqueue

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"
	"time"
)

type dqueue struct {
	path            string
	name            string
	maxBytesPerFile int
	writeFile       *os.File
	exitChan        chan int
	writeChan       chan []byte
	writeResp       chan error
	writePos        int64
	readChan        chan []byte
	writeFileNum    int64
	readLine        readPosition
	itemCounts      int64
	writeBuffer     bytes.Buffer
	sync.RWMutex
	needSync bool
	exitFlag int
}

func (d *dqueue) checkReadAble() bool {
	if d.readLine.readFileNum < d.writeFileNum || d.readLine.readPos < d.writePos {
		return true
	}
	return false
}

func (d *dqueue) fileName(fileNum int64) string {
	return fmt.Sprintf(path.Join(d.path, "%s.diskqueue.%06d.dat"), d.name, fileNum)
}
func (d *dqueue) ReadChan() <-chan []byte {
	return d.readChan
}
func (d *dqueue) metaDataFileName() string {
	return fmt.Sprintf(path.Join(d.path, "%s.meta.dat"), d.name)
}

func (d *dqueue) restoreFromMeta() error {
	fname := d.metaDataFileName()
	f, err := os.OpenFile(fname, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	var counts int64
	_, err = fmt.Fscanf(f, "%d\n%d,%d\n%d,%d\n",
		&counts,
		&d.writeFileNum, &d.writePos,
		&d.readLine.readFileNum, &d.readLine.readPos)
	if err != nil {
		return err
	}
	d.itemCounts = counts
	return nil
}

func (d *dqueue) saveMeta() error {
	tmpfile := fmt.Sprintf("%s/%s.%d.tmp", d.path, d.name, rand.Int())
	filename := d.metaDataFileName()
	f, err := os.OpenFile(tmpfile, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "%d\n%d,%d\n%d,%d\n", d.itemCounts, d.writeFileNum, d.writePos, d.readLine.readFileNum, d.readLine.readPos)
	if err != nil {
		f.Close()
		return err
	}
	f.Sync()
	f.Close()
	return os.Rename(tmpfile, filename)
}

func (d *dqueue) writeOne(data []byte) error {
	//TODO open file
	var err error
	if d.writeFile == nil {
		curFileName := d.fileName(d.writeFileNum)
		d.writeFile, err = os.OpenFile(curFileName, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		if d.writePos > 0 {
			_, err = d.writeFile.Seek(d.writePos, 0)
			if err != nil {
				d.writeFile.Close()
				d.writeFile = nil
				return err
			}
		}
	}
	//TODO filter data
	dataLen := int32(len(data))
	d.writeBuffer.Reset()
	err = binary.Write(&d.writeBuffer, binary.BigEndian, dataLen)
	if err != nil {
		return err
	}
	_, err = d.writeBuffer.Write(data)
	if err != nil {
		return err
	}
	_, err = d.writeFile.Write(d.writeBuffer.Bytes())
	if err != nil {
		d.writeFile.Close()
		d.writeFile = nil
		return err
	}
	d.writePos = d.writePos + 4 + int64(dataLen)
	d.itemCounts++

	if d.writePos > int64(d.maxBytesPerFile) {
		d.writeFileNum++
		d.sync()
		d.writeFile.Close()
		d.writeFile = nil
	}
	return err
}

func (d *dqueue) moveForward() error {
	d.itemCounts--
	return nil
}

func (d *dqueue) run() {
	var dataItem []byte
	//var r chan []byte
	for {
		select {
		case recv := <-d.writeChan:
			d.writeResp <- d.writeOne(recv)
		case d.readChan <- dataItem:
			d.moveForward()

		default:
			time.Sleep(1e9 * 3)
		}
		time.Sleep(1e9)
	}
}

func (d *dqueue) sync() error {
	if d.writeFile != nil {
		err := d.writeFile.Sync()
		if err != nil {
			d.writeFile.Close()
			d.writeFile = nil
			return err
		}
	}
	if err := d.saveMeta(); err != nil {
		return err
	}
	d.needSync = false
	return nil
}

func (d *dqueue) Put(data []byte) error {
	d.RLock()
	defer d.RUnlock()
	if d.exitFlag == 1 {
		return errors.New("exiting")
	}
	d.writeChan <- data
	return <-d.writeResp
}

func (d *dqueue) Close() error {
	d.RLock()
	defer d.RUnlock()
	d.exitFlag = 1
	close(d.exitChan)
	return d.sync()
}

type DB interface {
	Close() error
	Put(data []byte) error
}

func New(path, name string) DB {
	d := dqueue{
		path:            path,
		name:            name,
		maxBytesPerFile: 1 << 12,
		exitChan:        make(chan int, 0),
		writeChan:       make(chan []byte),
		writeResp:       make(chan error),
		readChan:        make(chan []byte),
	}
	err := d.restoreFromMeta()
	if err != nil {
		fmt.Println(err)
	}
	go d.run()
	return &d
}
