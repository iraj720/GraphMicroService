package data_handler

import (
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

type DiskDataHandler interface {
	ReadWriteHandler
	Mock(size int64) (string, error)
}
type diskDataHandler struct {
	f      os.File
	seeker int64
}

func NewDiskDataHandler(f os.File) DiskDataHandler {
	return &diskDataHandler{f: f, seeker: 0}
}

func (ddh *diskDataHandler) Read(gd *graphData) (n int, err error) {
	ddh.f.Seek(ddh.seeker, 0)
	var bytes []byte
	n, err = ddh.f.Read(bytes)
	if err != nil {
		return n, err
	}
	err = bson.Unmarshal(bytes, gd)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (ddh *diskDataHandler) Write(gd *graphData) (n int, err error) {
	_, err = ddh.f.Seek(0, 2)
	if err != nil {
		return 0, err
	}
	bsonm := bson.M{"ID": gd.ID, "Content": gd.Content, "TransferState": gd.TransferState}
	bytes, err := bson.Marshal(bsonm)
	if err != nil {
		return 0, err
	}
	return ddh.f.Write(bytes)
}

func (ddh *diskDataHandler) Close() error {
	return ddh.f.Close()
}

func (ddh *diskDataHandler) Size() int64 {
	stat, err := ddh.f.Stat()
	if err != nil {
		return -1
	}
	return stat.Size()
}
func (ddh *diskDataHandler) Mock(size int64) (string, error) {
	idGenerator := int64(0)
	for i := int64(0); i < size; i++ {
		// TODO : randomize size of content between 8-50000
		gd := &graphData{ID: idGenerator, Content: make([]byte, 500), TransferState: TransferState_INITIATED}
		ddh.Write(gd)
		idGenerator++
	}
	return ddh.f.Name(), nil
}
