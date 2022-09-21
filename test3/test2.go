package main

import(
	. "fmt"
	"regexp"
	"strings"
	"os"
)

func main() {
	Print(os.Args[1])
	Print(os.Args[2])
}

func a() {
	Println("abc")
	a := regexp.MustCompile(`${contextName}`)
	str := "this is ${contextName}\nthis is ${contextName}\nthis is ${contextName}"

	s := a.ReplaceAllString(str, "rep")
	s = strings.ReplaceAll(s, "${contextName}", "asdf")
	Println(s)
}
