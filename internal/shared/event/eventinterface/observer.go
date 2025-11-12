package eventinterface

type Observer[T any] interface {
	UpdateData(data T) error
}
