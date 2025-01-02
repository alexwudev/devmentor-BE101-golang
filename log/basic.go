package log

type Basic struct {
}

func (b *Basic) Show(loginfo string) {
	println(loginfo)
}
