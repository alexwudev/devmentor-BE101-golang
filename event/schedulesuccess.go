package event

type ScheduleSuccess struct {
	Basic
}

/*
func (e *ScheduleSuccess) Trigger(p person.PersonInterface) {
	e.SetNotify(notification.Email{})
	e.SetNotify(notification.Telegram{})

	for _, notify := range e.notifies {
		notify.Send(p, p.Speak(e.GetName()))
	}
}

func (e *ScheduleSuccess) GetName() string {
	return constants.ScheduleSuccess
}
*/
