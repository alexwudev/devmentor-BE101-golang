package log

type Base struct {
}

func (b *Base) Show(loginfo string) {
	println(loginfo)
}
