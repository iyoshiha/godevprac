package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"bufio"
	"io/ioutil"
	"regexp"
)

const targetPath = "a/b/c/d/e/target/"
// main jdbc
const pgDbms =  "DBMS=POSTGRES"
const pgDriver = "jdbc.driver=org.postgresql.Driver"
const pgUrl = "jdbc.url=jdbc:postgresql://"
const sqlDbms = "DBMS=SQLSERVER"
const sqlDriver = "jdbc.driver=com.microsoft.sqlserver.jdbc.SQLServerDriver"
const sqlUrl = "jdbc.url=jdbc:sqlserver://"
const jdbcName = "jdbc.username="
const jdbcPass = "jdbc.password="
// sub jdbc
const subpgDbms =  "DBMS.sub=POSTGRES"
const subpgDriver = "jdbc.sub.driver=org.postgresql.Driver"
const subpgUrl = "jdbc.sub.url=jdbc:postgresql://"
const subsqlDbms = "DBMS.sub=SQLSERVER"
const subsqlDriver = "jdbc.sub.driver=com.microsoft.sqlserver.jdbc.SQLServerDriver"
const subsqlUrl = "jdbc.sub.url=jdbc:sqlserver://"
const subjdbcName = "jdbc.sub.username="
const subjdbcPass = "jdbc.sub.password="

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
	strs := getValueFromConfig()
	allInfo := createAllInfo(strs[5:], strs[:5])

	file, err := os.Create("resource.properties")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	templateByte, err := ioutil.ReadFile("resource.properties.template")
	if err != nil {
		log.Fatal(err)
	}

	templateStr := string(templateByte) 

	templateStr = templateStr +
	allInfo.mainJdbc.Dbms +"\n"+
	allInfo.mainJdbc.Driver+"\n"+
	allInfo.mainJdbc.Domain+"\n"+
	allInfo.mainJdbc.Username+"\n"+
	allInfo.mainJdbc.Password+"\n"+
	allInfo.subJdbc.Dbms +"\n"+
	allInfo.subJdbc.Driver+"\n"+
	allInfo.subJdbc.Domain+"\n"+
	allInfo.subJdbc.Username+"\n"+
	allInfo.subJdbc.Password+"\n"

	_, err = file.WriteString(templateStr)
	if err != nil {
		log.Fatal(err)
	}
}

func printAllInfoContents(allInfo AllInfo){
	fmt.Println(allInfo.pathInfo)
	fmt.Println(allInfo.mainJdbc)
	fmt.Println(allInfo.subJdbc)
}

func getValueFromConfig() []string{
	config, err := os.Open("test_config.properties")
	var properties []string
	regEqual := regexp.MustCompile(`=`)

	if err != nil {
		panic(err)
	}

	defer config.Close()


	scanner := bufio.NewScanner(config)

	// for scanner.Scan() {
	// 	scanner.Text()
	// 	firstChar := ""
	// 	if "#" == firstChar {
	// 		continue
	// 	}
	// }

	// 先頭１文字ががシャープ（コメント）かどうか確認
	// propertyは、properties sliceに格納
	for str := ""; scanner.Scan(); {
		str = scanner.Text()
		for _, c :=  range str{
			if c == '#' {
				break
			} 
			strSplited := regEqual.Split(str,-1)
			properties = append(properties, strSplited[1])
			break
		}
	}
	return properties
}

// func setPropaties(dbType string, sub bool) JdbcInfo{

// 	args, dbType := commandLineArgs()
// 	validateDbType(dbType)
// 	if len(args) > 0 {
// 		log.Fatal("\"-db=\"でDB種類を指定してください")
// 	}

// 	strs := getValueFromConfig()
// 	fmt.Println(strs)
// 	pathInfo = strs[:5]
// 	dbinfo = strs[5:]
	// return jdbcInfo
// }

type PathInfo struct {
	gitPath string
	wesPath string
	esm_log string
}

type AllInfo struct {
	mainJdbc JdbcInfo
	subJdbc JdbcInfo
	pathInfo PathInfo
}

func createAllInfo(jdbcStr, pathStr []string) AllInfo{

	var dbms string
	var driver string
	var url string 
	var port string
	dbName := "" 
	
	if dbtype := jdbcStr[0]; "pg" == dbtype {
		dbms = pgDbms
		driver = pgDriver
		url = pgUrl+jdbcStr[2]+""
		dbName = "/" + jdbcStr[2]
		port = "5432"
	} else if "sqlserver" == dbtype {
		dbms = sqlDbms
		driver = sqlDriver
		url = sqlUrl
		port = "1433" 
	}

	jdbcInfo := JdbcInfo{
		Dbms: dbms,
		Driver: driver,
		Domain: url+jdbcStr[1]+":"+port+dbName, 
		Username: jdbcName + jdbcStr[2],
		Password: jdbcPass + jdbcStr[3],
	}

	if dbtype := jdbcStr[0]; "pg" == dbtype {
		dbms = subpgDbms
		driver = subpgDriver
		url = subpgUrl+jdbcStr[2]+""
		dbName = "/" + jdbcStr[2]
		port = "5432"
	} else if "sqlserver" == dbtype {
		dbms = subsqlDbms
		driver = subsqlDriver
		url = subsqlUrl
		port = "1433" 
	}

	subJdbcInfo := JdbcInfo{
		Dbms: dbms,
		Driver: driver,
		Domain: url+jdbcStr[1]+":"+port+dbName, 
		Username: subjdbcName + jdbcStr[2]+"_sub",
		Password: subjdbcPass + jdbcStr[3]+"_sub",
	}

	pathInfo := PathInfo{
		gitPath:pathStr[0],
		wesPath:pathStr[0] + "/repo" + "/remix/esm_war",
		esm_log:pathStr[4],
	}
	allInfo := AllInfo{jdbcInfo, subJdbcInfo, pathInfo}
	return allInfo
}

func validateAll() {
	// args validation
	args, dbType := commandLineArgs()
	validateDbType(dbType)
	if len(args) > 0 {
		log.Fatal("\"-db=\"でDB種類を指定してください")
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
	args = flag.Args()

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