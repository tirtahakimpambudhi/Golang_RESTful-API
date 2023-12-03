package config

import (
	"fmt"
	"os"
	"path/filepath"
	"restful_api/helper"
	"github.com/joho/godotenv"
)
var workdir , _  = os.Getwd()
// var BASEDIR = filepath.Join(workdir,"..")//jika untuk test
var BASEDIR = filepath.Join(workdir) //jika main.go
var DBHOST, DBPORT, DBUSER, DBPASS, DBNAME , TBNAME , DIALECT, CONNECT , HOST , PORT , ADDR string

func init() {
	err := godotenv.Load(filepath.Join(BASEDIR,".env"))
	helper.PanicIFError(err)

	DBHOST = os.Getenv("dbhost")
	DBPORT = os.Getenv("dbport")
	DBUSER = os.Getenv("dbuser")
	DBPASS = os.Getenv("dbpass")
	DIALECT = os.Getenv("dialect")
	DBNAME = os.Getenv("dbname")
	TBNAME = os.Getenv("tbname")
	HOST = os.Getenv("host")
	PORT = os.Getenv("port")
	ADDR = fmt.Sprintf("%s:%s",HOST,PORT)
	CONNECT = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",DBUSER,DBPASS,DBHOST,DBPORT,DBNAME)
}
