package event

import (
	eventinterface "be-mission-management/src/shared/event/interface"
	"sync"
)

type BaseObservable[T any] struct {
	observers map[eventinterface.Observer[T]]struct{}
	lock      sync.Mutex
}

func NewBaseObservable[T any]() *BaseObservable[T] {
	return &BaseObservable[T]{
		observers: make(map[eventinterface.Observer[T]]struct{}),
	}
}

func (o *BaseObservable[T]) Subscribe(observer eventinterface.Observer[T]) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.observers[observer] = struct{}{}
}

func (o *BaseObservable[T]) Unsubscribe(observer eventinterface.Observer[T]) {
	o.lock.Lock()
	defer o.lock.Unlock()
	delete(o.observers, observer)
}

func (o *BaseObservable[T]) Notify(data T) {
	o.lock.Lock()
	defer o.lock.Unlock()
	for obs := range o.observers {
		go obs.UpdateData(data)
	}
}
