package enums

/*
all of the files in this folder belong to the enums package, denoting enums that are useful for
something or other throughout the rest of the wrapper.
*/
type Language int

const (
	English Language = iota + 1
	Chinese
	Japanese
	Korean
)

func (l Language) String() string {
	return [...]string{"en", "cn", "jp", "kr"}[l-1]
}

func (l Language) EnumIndex() int {
	return int(l)
}
