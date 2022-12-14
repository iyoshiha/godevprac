package main

import(
	. "fmt"
	"regexp"
	"strings"
	"os"
	"log"
)

const path = "/opt/remix/work/resource.properties"

func main() {
	b()
}

func b() {
	file, err := os.Create(path)
	if nil != err {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("this is wrigting\n"+
	"ok man\r"+
	"nice work\r\n")
	if err != nil {
		log.Fatal(err)
	}
}

func a() {
	Println("abc")
	a := regexp.MustCompile(`${contextName}`)
	str := "this is ${contextName}\nthis is ${contextName}\nthis is ${contextName}"

	s := a.ReplaceAllString(str, "rep")
	s = strings.ReplaceAll(s, "${contextName}", "asdf")
	Println(s)
}
