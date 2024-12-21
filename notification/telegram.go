package notification

import (
	"be101/person"
	"fmt"
)

type Telegram struct {
	PhoneNumber string
}

func (telegram Telegram) Send(p person.PersonInterface, s string) {
	fmt.Println(s + " by Sending Telegram notification! ")
}
