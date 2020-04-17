package diskqueue

import "os"

type readPosition struct {
	readPos     int64
	readFileNum int64
	readFile    *os.File
	d           *dqueue
}

func (r *readPosition) readOne() (data []byte, err error) {
	if r.readFile == nil {
		curFileName := r.d.fileName(r.readFileNum)
		r.readFile, err = os.OpenFile(curFileName, os.O_RDONLY, 0600)
		if err != nil {
			return
		}
		if r.readPos > 0 {
			_, err = r.readFile.Seek(r.readPos, 0)
			if err != nil {
				r.readFile.Close()
				r.readFile = nil
				return
			}
		}
	}
	return
}
