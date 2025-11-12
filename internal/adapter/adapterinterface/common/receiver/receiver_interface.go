package commonreceiverinterface

type ReceiverInterface interface {
	Start() error
	Stop() error
}
