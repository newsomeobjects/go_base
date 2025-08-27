package main

import "fmt"

func main() {
	var a = "mum"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	var s = "hello"
	switch { //注意，这种情况的 switch 后面不再需要跟判断变量。
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}
