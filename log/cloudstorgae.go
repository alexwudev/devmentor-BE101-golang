package log

import "fmt"

type CloudStorage struct {
	Base
}

func (c *CloudStorage) Show(loginfo string) {
	fmt.Println(loginfo)
}
