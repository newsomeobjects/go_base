package pk1

import (
	"fmt"

	_ "go_base/base/GoInitOrder/pk2"
)

const PkgName string = "pk1"

var PkgNameVar string = getPkgName()

func init() {
	fmt.Println("pkg1 init method invoked")
}

func getPkgName() string {
	fmt.Println(PkgName)
	fmt.Println("pkg1.PkgNameVar has been initialized")
	return "abc"
}
