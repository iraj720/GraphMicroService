package data_handler

import (
	"math/rand"
	"sync"
)

type memoryDataHandler struct {
	data   []*GraphData
	_lock  sync.RWMutex
	seeker int64
}

func NewMemoryDataHandler(withMockedData bool, mockDataSize int64) ReadWriteHandler {
	mdh := &memoryDataHandler{_lock: sync.RWMutex{}, seeker: 0, data: make([]*GraphData, mockDataSize)}
	if withMockedData {
		mdh.mock(mockDataSize)
	}
	return mdh
}

func (mdh *memoryDataHandler) Read(gd *GraphData) (n int, err error) {
	mdh._lock.RLock()
	defer mdh._lock.RUnlock()
	gd.Content = mdh.data[mdh.seeker].Content
	gd.ID = mdh.data[mdh.seeker].ID
	gd.TransferState = mdh.data[mdh.seeker].TransferState
	mdh.seeker++
	return len(gd.Content), nil
}

// TODO: handle unused data overflow
func (mdh *memoryDataHandler) Write(gd *GraphData) (n int, err error) {
	mdh._lock.Lock()
	defer mdh._lock.Unlock()
	mdh.data = append(mdh.data, gd)
	_ = mdh.data[gd.ID]
	return len(mdh.data), nil
}

func (mdh *memoryDataHandler) Close() error {
	_ = mdh.data
	return nil
}

func (mdh *memoryDataHandler) Size() int64 {
	return int64(len(mdh.data))
}

func (mdh *memoryDataHandler) mock(size int64) {
	idGenerator := int64(0)
	max := 50000
	min := 8
	for i := int64(0); i < size; i++ {
		// TODO : randomize size of content between 8-50000
		number := rand.Int31n(int32(max-min)) + int32(min)
		// 8 B - 5 KB data
		b := make([]byte, number)
		mdh.data[i] = &GraphData{ID: idGenerator, Content: b, TransferState: TransferState_INITIATED}
		idGenerator++
	}
}
