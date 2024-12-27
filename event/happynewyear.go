package event

import (
	"be101/constants"
	"be101/notification"
	"be101/person"
)

type HappyNewYear struct {
	Basic
}

func (e *HappyNewYear) Trigger(p person.PersonInterface) {
	e.SetNotify(notification.Line{})

	for _, notify := range e.notifies {
		notify.Send(p, p.Speak(e.GetName()))
	}

}

func (e *HappyNewYear) GetName() string {
	return constants.HappyNewYear
}
