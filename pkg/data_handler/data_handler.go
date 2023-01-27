package data_handler

type DataHandler interface {
	GraceFullShutdown()
	// it will read the ds from start to end no matter if anything fails
	ReadData() (*GraphData, error)

	// it will write data to end of the ds
	WriteData(gd *GraphData) error

	// if data_handler has ran in withMockData mode its len is constant else it is 0 and will grow based on your writes
	Size() int64
}

// conterpart of io.ReadWriteCloser
type ReadWriteHandler interface {
	Read(*GraphData) (int, error)
	Write(*GraphData) (int, error)
	Close() error
	Size() int64
}
type dataHandler struct {
	rwc ReadWriteHandler
}

func NewDataHandler(rwc ReadWriteHandler) DataHandler {
	return &dataHandler{rwc: rwc}
}

func (dh *dataHandler) GraceFullShutdown() {
	// for i := int64(0); i < dh.rwc.Size(); i++ {
	// 	err := dh.UpdateData(&dh.data[i])
	// 	if err != nil {
	// 		// failed to handle failed data
	// 	}
	// }
}

func (dh *dataHandler) ReadData() (*GraphData, error) {
	gd := &GraphData{}
	_, err := dh.rwc.Read(gd)
	if err != nil {
		return nil, err
	}
	return gd, nil
}

func (dh *dataHandler) WriteData(gd *GraphData) error {
	_, err := dh.rwc.Write(gd)
	if err != nil {
		return err
	}
	return nil
}

func (dh *dataHandler) Size() int64 {
	return dh.rwc.Size()
}
