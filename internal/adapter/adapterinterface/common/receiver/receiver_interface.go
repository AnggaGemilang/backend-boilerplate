package commonreceivergateway

type ReceiverInterface interface {
	Start() error
	Stop() error
}
