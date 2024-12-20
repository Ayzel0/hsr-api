package enums

/*
all of the files in this folder belong to the enums package, denoting enums that are useful for
something or other throughout the rest of the wrapper.
*/

type Element int

const (
	Physical Element = iota + 1
	Fire
	Ice
	Lightning
	Wind
	Quantum
	Imaginary
)

func (e Element) String() string {
	return [...]string{"Physical", "Fire", "Ice", "Lightning", "Wind", "Quantum", "Imaginary"}[e-1]
}

func (e Element) EnumIndex() int {
	return int(e)
}
