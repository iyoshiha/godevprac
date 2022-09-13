package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"bufio"

)

const targetPath = "a/b/c/d/e/target/"

type JdbcInfo struct {
	Dbms		string
	Driver		string
	Domain		string
	Username	string
	Password	string
}

func main() {

// dbtype=
// domain=
// username=
// password=
// contextName=



}

func examConfigFile() {
	config, err =: os.Open(test_config.properties)

	if err != nil {
		panic(err)
	}

	defer config.Close()


	scanner := bufio.NewScanner(config)

	for scanner.Scan() {
		fmt.Println(scanner.Test())
	}

}


func setPropaties(dbType string, sub bool) JdbcInfo{

	args, dbType := commandLineArgs()
	validateDbType(dbType)
	if args.size() > 0 {
		log.fatal("\"-db=\"でDB種類を指定してください")
	}

	os.Open()

	jdbcInfo := {
		Dbms:
		Driver:
		Domain:
		Username:
		Password:
		
	}

	return 
}

func validateAll() {
	// args validation
	args, dbType := commandLineArgs()
	validateDbType(dbType)
	if args.size() > 0 {
		log.fatal("\"-db=\"でDB種類を指定してください")
	}

}



func createAndWriteEmptyfile() {
	file, err := os.Create("empty.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
	fmt.Println("file created")


    words := []string{"sky", "falcon", "rock", "hawk"}

    for _, word := range words {

        _, err := file.WriteString(word + "\n")

        if err != nil {
            log.Fatal(err)
        }
    }

}

func validateDbType(dbType string) {
	switch dbType {
	case "sqlserver":
	case "pg", "postgresql", "postgres":
	default:
		log.Fatalln("databaseの種類が不正です:", dbType, "sqlサーバーの場合引数はいりません。タイプミスの可能性")
	}
	fmt.Println("database type:", dbType)
}

func commandLineArgs() (args []string, dbType string){
	// flagで指定して与えられた値はArgsで参照不可
	f := flag.String("db", "sqlserver", "sqlserver or postgres")
	

	// これ以降フラグの設定など不可能
	//　設定する場合、この行より上に書く
	flag.Parse()
	str := flag.Args()

	dbType = *f
	fmt.Println("選択されたDBの種類:", dbType)
	return args, dbType
}


func late() {
	defer fmt.Println("defer")
	fmt.Println("late is called")
}
func secondLate() {
	defer fmt.Println("second defer")
	fmt.Println("second late is called")
}

func writeFile(contents []string) {


}