package Observable

import (
	. "ObserverPattern/Types"
)

type IPhoneObservable struct {
	StockUnit int
	Observers []IAlert
}

func (iPhoneObservable *IPhoneObservable) NewIPhoneObservable() *IPhoneObservable {
	return &IPhoneObservable{
		StockUnit: iPhoneObservable.StockUnit,
		Observers: iPhoneObservable.Observers,
	}
}

func (iPhoneObservable *IPhoneObservable) Add(observers IAlert) {
	iPhoneObservable.Observers = append(iPhoneObservable.Observers, observers)
}

func (iPhoneObservable *IPhoneObservable) Remove(observers IAlert) {
	for i, observer := range iPhoneObservable.Observers {
		if observer == observers {
			iPhoneObservable.Observers = append(iPhoneObservable.Observers[:i], iPhoneObservable.Observers[i+1:]...)
		}
	}
}

func (iPhoneObservable *IPhoneObservable) Notify() {
	for _, observer := range iPhoneObservable.Observers {
		observer.Update()
	}
}

func (iPhoneObservable *IPhoneObservable) SetData(data int) {
	if iPhoneObservable.StockUnit == 0 {
		iPhoneObservable.Notify()
	}
	iPhoneObservable.StockUnit = data
}

func (iPhoneObservable *IPhoneObservable) GetData() int {
	return iPhoneObservable.StockUnit
}
