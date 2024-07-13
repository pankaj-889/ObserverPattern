package Observer

import (
	. "ObserverPattern/Types"
	"fmt"
)

type EmailAlert struct {
	Email      string
	Observable IStockObservable
}

func NewEmailAlert(email string, observable IStockObservable) *EmailAlert {
	return &EmailAlert{
		Email:      email,
		Observable: observable,
	}
}

func (e *EmailAlert) Update() {
	e.SendEmail(e.Email)
}

func (e *EmailAlert) SendEmail(message string) {
	fmt.Println("Email Alert for", message)
}
