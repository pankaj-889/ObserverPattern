package main

import (
	. "ObserverPattern/Observable"
	. "ObserverPattern/Observer"
	. "ObserverPattern/Types"
)

func main() {

	phone := IPhoneObservable{
		StockUnit: 0,
		Observers: []IAlert{},
	}
	Iphone := phone.NewIPhoneObservable()
	MobileAlert := NewMobileAlert("Mobile", Iphone)
	EmailAlert := NewEmailAlert("Email", Iphone)

	Iphone.Add(MobileAlert)
	Iphone.Add(EmailAlert)

	Iphone.SetData(10)
	Iphone.SetData(0)
	Iphone.SetData(20)
}
