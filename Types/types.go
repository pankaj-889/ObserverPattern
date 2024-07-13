package Types

type IAlert interface {
	Update()
}

type IStockObservable interface {
	Add(observers IAlert)
	Remove(observers IAlert)
	SetData(data int)
	GetData() int
	Notify()
}
