package data_handler

import "sync"

type MemoryDataHandler interface {
	ReadWriteHandler
	Mock(size int64) ([]graphData, error)
}
type memoryDataHandler struct {
	data   []graphData
	_lock  sync.RWMutex
	seeker int64
}

func NewMemoryDataHandler() MemoryDataHandler {
	return &memoryDataHandler{_lock: sync.RWMutex{}, seeker: 0}
}

func (mdh *memoryDataHandler) Read(gd *graphData) (n int, err error) {
	mdh._lock.RLock()
	defer mdh._lock.RUnlock()
	gd = &mdh.data[mdh.seeker]
	mdh.seeker++
	return len(gd.Content), nil
}

// TODO: handle unused data overflow
func (mdh *memoryDataHandler) Write(gd *graphData) (n int, err error) {
	mdh._lock.Lock()
	defer mdh._lock.Unlock()
	mdh.data = append(mdh.data, *gd)
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

func (mdh *memoryDataHandler) Mock(size int64) ([]graphData, error) {
	idGenerator := int64(0)
	gds := make([]graphData, size)
	for i := int64(0); i < size; i++ {
		// TODO : randomize size of content between 8-50000
		gds[i] = graphData{ID: idGenerator, Content: make([]byte, 500), TransferState: TransferState_INITIATED}
		idGenerator++
	}
	return gds, nil
}
