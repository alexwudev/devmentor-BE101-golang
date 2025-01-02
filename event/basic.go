package event

import (
	"be101/constants"
	"be101/notification"
	"be101/person"
)

type Basic struct {
	notifies  []notification.NotificationInterface
	eventName string
}

func (b *Basic) SetNotify(n notification.NotificationInterface) {
	b.notifies = append(b.notifies, n)
}

func (b *Basic) Trigger(p person.PersonInterface) {
	switch b.eventName {
	case constants.CancelClasses:
		b.SetNotify(notification.Email{})
		b.SetNotify(notification.Telegram{})
	case constants.HappyNewYear:
		b.SetNotify(notification.Line{})
	case constants.RegisterSuccess:
		b.SetNotify(notification.Email{})
		b.SetNotify(notification.Sms{})
	case constants.ScheduleSuccess:
		b.SetNotify(notification.Email{})
		b.SetNotify(notification.Telegram{})
	}
	for _, notify := range b.notifies {
		notify.Send(p, p.Speak(b.eventName))
	}
}

func (b *Basic) SetEventName(eventName string) {
	b.eventName = eventName
}

/*
	func (b *Basic) GetName() string {
		return constants.CancelClasses
	}
*/
