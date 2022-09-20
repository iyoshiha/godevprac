package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"bufio"
	"regexp"
	"strings"
	"io/ioutil"
)

const remixPath = "/opt/remix/"
const confPath= remixPath+"remix.conf"
const targetPath = remixPath+"work/"
const templatePath = remixPath+"template/"
const resourcePropertiesTemplateFile = templatePath+"resource.properties.template"
const resourceProperties = "remix/esm_war/WEB-INF/src/jp/co/softbrain/wes/"
// main jdbc
const pgDbms =  "DBMS=POSTGRES"
const pgDriver = "jdbc.driver=org.postgresql.Driver"
const pgUrl = "jdbc.url=jdbc:postgresql://"
const sqlDbms = "DBMS=SQLSERVER"
const sqlDriver = "jdbc.driver=com.microsoft.sqlserver.jdbc.SQLServerDriver"
const sqlUrl = "jdbc.url=jdbc:sqlserver://"
const jdbcName = "jdbc.username="
const jdbcPass = "jdbc.password="
const pgPort = "5432"
const sqlPort = "1433"
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

type confItems struct {
	repository string
	branch string
	dbtype	string
	domain	string
	username	string
	contextName	string
	password	string
	subUsername	string
	subPassword	string
	dbname	string
}


func main() {
	allInfo := createAllInfo(setConfItems(readConfItems(confPath)))
	fmt.Println(writeRecourceProperties(allInfo))
}

/*
func temporary() {
	strs := getValueFromConfig()
	allInfo := createAllInfo(strs[5:], strs[:5])

	file, err := os.Create(targetPath+"resource.properties")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	templateByte, err := ioutil.ReadFile(templatePath+"resource.properties.template")
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
*/
// necesary j;lkj;kljkljklj
func readConfItems (filePath string) []string {
	configFile, err := os.Open(filePath)
	var properties []string
	regEqual := regexp.MustCompile(`=`)

	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)

	// 先頭１文字ががシャープ（コメント）かどうか確認
	// propertyは、properties sliceに格納
	for str := ""; scanner.Scan(); {
		str = strings.TrimSpace(scanner.Text())
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

func setConfItems(properties []string) (conf confItems) {
	subFlag := false
	conf.repository	= properties[0]
	conf.branch	= properties[1]
	conf.dbtype	= properties[2]
	conf.domain	= properties[3]
	conf.username	= properties[4]
	conf.contextName	= properties[5]
	// ********  optional items ***********
	conf.password	= getOptionalConf(properties[6], conf.username, subFlag)
	subFlag = true
	conf.subUsername	= getOptionalConf(properties[7], conf.username, subFlag)
	conf.subPassword	= getOptionalConf(properties[8], conf.username, subFlag)
	subFlag = false
	conf.dbname	= getOptionalConf(properties[9], conf.username, subFlag)

	return conf
}

func getOptionalConf(option, username string, subFlag bool) (confVal string){
	if option != "" {
		confVal = option
	} else {
		confVal = username
	}
	if subFlag {
		confVal += "_sub"
	}
	return confVal
}

func writeRecourceProperties(allInfo AllInfo) string{
	templateByte, err := ioutil.ReadFile(resourcePropertiesTemplateFile)
	if err != nil {
		log.Fatal(err)
	}

	templateStr := string(templateByte)

	templateStr = allInfo.mainJdbc.Dbms +"\n"+
	allInfo.mainJdbc.Driver+"\n"+
	allInfo.mainJdbc.Domain+"\n"+
	allInfo.mainJdbc.Username+"\n"+
	allInfo.mainJdbc.Password+"\n"+
	allInfo.subJdbc.Dbms +"\n"+
	allInfo.subJdbc.Driver+"\n"+
	allInfo.subJdbc.Domain+"\n"+
	allInfo.subJdbc.Username+"\n"+
	allInfo.subJdbc.Password+"\n"+
	templateStr

//	_, err = file.WriteString(templateStr)
//	if err != nil {
//		log.Fatal(err)
//	}

	return templateStr
}

type AllInfo struct {
	mainJdbc JdbcInfo
	subJdbc JdbcInfo
}

func createAllInfo(conf confItems) AllInfo{

	var dbms string
	var driver string
	var url string
	var port string
	dbName := ""

	// main jdbc info is being filled
	if dbtype := conf.dbtype; "pg" == dbtype {
		dbms = pgDbms
		driver = pgDriver
		url = pgUrl
		dbName = "/" + conf.dbname
		port = pgPort
	} else if "sqlserver" == dbtype {
		dbms = sqlDbms
		driver = sqlDriver
		url = sqlUrl
		port = sqlPort
	}

	jdbcInfo := JdbcInfo{
		Dbms: dbms,
		Driver: driver,
		Domain: url+conf.domain+":"+port+dbName,
		Username: jdbcName + conf.username,
		Password: jdbcPass + conf.password,
	}

	// sub jdbc info is being filled
	if conf.dbtype == "pg"{
		dbms = subpgDbms
		driver = subpgDriver
		url = subpgUrl
		dbName = "/" + conf.dbname
		port = "5432"
	} else if "sqlserver" == conf.dbtype {
		dbms = subsqlDbms
		driver = subsqlDriver
		url = subsqlUrl
		port = "1433"
	}

	subJdbcInfo := JdbcInfo{
		Dbms: dbms,
		Driver: driver,
		Domain: url+conf.domain+":"+port+dbName,
		Username: subjdbcName + conf.subUsername,
		Password: subjdbcPass + conf.subPassword,
	}

	allInfo := AllInfo{jdbcInfo, subJdbcInfo,}
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
