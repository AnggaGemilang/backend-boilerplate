package eventinterface

type Observable[T any] interface {
	Subscribe(observer Observer[T])
	Unsubscribe(observer Observer[T])
	Notify(data T)
}
