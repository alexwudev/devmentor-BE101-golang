package event

import (
	"be101/notification"
)

type Basic struct {
	notifies []notification.NotificationInterface
}

func (b *Basic) SetNotify(n notification.NotificationInterface) {
	b.notifies = append(b.notifies, n)
}
