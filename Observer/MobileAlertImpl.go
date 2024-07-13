package Observer

import (
	. "ObserverPattern/Types"
	"fmt"
)

type MobileAlert struct {
	Name       string
	Observable IStockObservable
}

func NewMobileAlert(name string, observable IStockObservable) *MobileAlert {
	return &MobileAlert{
		Name:       name,
		Observable: observable,
	}
}

func (m *MobileAlert) Update() {
	m.SendNotification(m.Name)
}

func (m *MobileAlert) SendNotification(message string) {
	fmt.Println("Mobile Alert for", message)
}
