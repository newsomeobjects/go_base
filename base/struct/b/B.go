package b

type B struct {
	b     string
	bytes [2]byte
	int
	C
}

type C struct {
	b     string
	bytes [2]byte
	int
}

func (b *B) SetB(b1 string) {
	b.b = b1
}

func NewB() B {
	return B{b: "bbb", bytes: [2]byte{'b', 'b'}, int: 2, C: C{b: "ccc", bytes: [2]byte{'b', 'b'}}}
}
