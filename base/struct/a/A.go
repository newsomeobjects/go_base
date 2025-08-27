package a

type A struct {
	a     string
	bytes [2]byte
	int
}

func NewA() A {
	return A{a: "aaa", bytes: [2]byte{'a', 'a'}, int: 0}
}

func (a1 A) SetA(a string) { a1.a = a }

//func (a1 A) GetA() string  { return a1.a }

func (a1 *A) SetPA(a string) { a1.a = a }
