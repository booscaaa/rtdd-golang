package hello

import "fmt"

func PrintHello(name string) string {
	a := 1
	b := 2

	if a == b {
		fmt.Println("É igual")
	}
	return "Hello, " + name
}
