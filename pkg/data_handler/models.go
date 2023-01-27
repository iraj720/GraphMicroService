package data_handler

type TransferState int

const (
	TransferState_DEFAULT   TransferState = 2
	TransferState_INITIATED TransferState = 3
	TransferState_SENT      TransferState = 0
	TransferState_FAILED    TransferState = 1
)

type GraphData struct {
	ID            int64
	Content       []byte
	TransferState TransferState
}
