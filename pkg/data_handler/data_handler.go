package data_handler

type DataHandler interface {
}

// conterpart of io.ReadWriteCloser
type ReadWriteHandler interface {
	Read(*graphData) (int, error)
	Write(*graphData) (int, error)
	Close() error
	Size() int64
}
type dataHandler struct {
	rwc ReadWriteHandler
}

func NewDataHandler(rwc ReadWriteHandler) DataHandler {
	return dataHandler{rwc: rwc}
}

func (dh *dataHandler) GraceFullShutdown() {
	// for i := int64(0); i < dh.rwc.Size(); i++ {
	// 	err := dh.UpdateData(&dh.data[i])
	// 	if err != nil {
	// 		// failed to handle failed data
	// 	}
	// }
}

func (dh *dataHandler) UpdateData(gd *graphData) error {
	_, err := dh.rwc.Write(gd)
	if err != nil {
		return err
	}
	return nil
}
