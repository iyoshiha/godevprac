package main

import(
	. "fmt"
	"os"
	"regexp"
	"bufio"
	"strings"
)

const remixConfPath="remix.conf"

func main() {
	Println(getValueFromConfig())
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
			} 
// else if  c == ' ' {
//				break
//			}
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
