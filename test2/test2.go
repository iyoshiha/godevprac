package main

import(
	. "fmt"
	"os"
	"regexp"
	"bufio"
	"strings"
)

const remixConfPath="remix.conf"

type bc struct {
	a string
	b string
	c string
}

func main() {
	str := " "
	for _, v := range str{
		Println("++++++++++++")
		Println(string(v))
		Println("++++++++++++")
	}

}


func pracstruct () {
	abc := bc{a: "a", b:"",}
	Println("start")
	Println(len(abc.c))
	Println(abc.c)
	Println("end")
}

func j() {
	str := strings.TrimSpace("      ")
	Println(len(str))
}
func sc() {
	Println("")
	regEqual := regexp.MustCompile(`=`)
	f, _ := os.Open("txt")
	scanner := bufio.NewScanner(f)

	// 先頭１文字ががシャープ（コメント）かどうか確認
	// propertyは、properties sliceに格納
	for str := ""; scanner.Scan(); {
		str = strings.TrimSpace(scanner.Text())
		strSplited := regEqual.Split(str,-1)
		Println("===========")
		Println(len(strSplited))
		Println(strSplited[0])
		Println("@@@@@@@@@@@")
	}

}

func aa() {
	str := []string{"abc=123", "", "def=", "", " "}
	regEqual := regexp.MustCompile(`=`)

	for _, v := range str {
		strSplited := regEqual.Split(v,-1)
		Println("===========")
		Println(strSplited)
		Println(len(strSplited))
		Println(strSplited[0])
		Println("@@@@@@@@@@@")
	}
}

func getValueFromConfig() []string{
	config, err := os.Open(remixConfPath)
	var properties []string
	regEqual := regexp.MustCompile(`=`)

	if err != nil {
		panic(err)
	}

	defer config.Close()


	scanner := bufio.NewScanner(config)

	// 先頭１文字ががシャープ（コメント）かどうか確認
	// propertyは、properties sliceに格納
	for str := ""; scanner.Scan(); {
		str = strings.TrimSpace(scanner.Text())
		println(str)
		for _, c :=  range str{
			if c == '#' {
				break
			} else if  c == ' ' {
				break
			}
			strSplited := regEqual.Split(str,-1)
		println(strSplited[0])
		println(strSplited)
			properties = append(properties, strSplited[1])
			break
		}
	}
	println(len(properties))
	return properties
}
