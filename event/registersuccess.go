package event

import (
	"be101/constants"
	"be101/notification"
	"be101/person"
)

type RegisterSuccess struct {
	Basic
}

func (e *RegisterSuccess) Trigger(p person.PersonInterface) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Sms{})

	for _, notify := range e.notifies {
		notify.Send(p, p.Speak(e.GetName()))
	}
}

func (e *RegisterSuccess) GetName() string {
	return constants.RegisterSuccess
}
