package main

import(
	. "fmt"
)

type st struct {
	a string
	b string
	c string
}
func main() {
	stval := st{"a", "d", "d"}
	for _, v := range stval {
		Println(v)

	}
	
}