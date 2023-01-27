package data_handler

import (
	"math/rand"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

type diskDataHandler struct {
	f      os.File
	seeker int64
}

func NewDiskDataHandler(f os.File, withMockedData bool, mockDataSize int64) ReadWriteHandler {
	ddh := &diskDataHandler{f: f, seeker: 0}
	if withMockedData {
		ddh.mock(mockDataSize)
	}
	return ddh
}

func (ddh *diskDataHandler) Read(gd *GraphData) (n int, err error) {
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

func (ddh *diskDataHandler) Write(gd *GraphData) (n int, err error) {
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
func (ddh *diskDataHandler) mock(size int64) {
	idGenerator := int64(0)
	max := 50000
	min := 8
	for i := int64(0); i < size; i++ {
		// TODO : randomize size of content between 8-50000
		number := rand.Int31n(int32(max-min)) + int32(min)
		// 8 B - 5 KB data
		b := make([]byte, number)
		gd := &GraphData{ID: idGenerator, Content: b, TransferState: TransferState_INITIATED}
		ddh.Write(gd)
		idGenerator++
	}
}
