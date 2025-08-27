package main

import "fmt"

type Supplier interface {
	Get() string
}

type DigitSupplier struct {
	value int
}

func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

func main() {
	var a Supplier = &DigitSupplier{value: 1}
	fmt.Println(a.Get())

	b, ok := a.(*DigitSupplier)
	fmt.Println(b, ok)
}
