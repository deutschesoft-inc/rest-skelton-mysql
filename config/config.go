package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	ent "main/entities"
	"os"
)

var (
	DB                     *sql.DB
	jsonFile               *os.File
	err                    error
	Secret, DevMode, IOURL string
	Secured                bool
	Auth                   ent.NetGsmAuth
)

type Configs struct {
	Dbhost string `json:"dbhost"`
	Dbport int    `json:"dbport"`
	Dbuser string `json:"dbuser"`
	Dbpass string `json:"dbpass"`
	Dbname string `json:"dbname"`
	Sslsup string `json:"sslsup"`
	Secret string `json:"secret"`
}

type Secrets struct {
	SecretKey string
}

func init() {
	jsonFile, err = os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf Configs
	json.Unmarshal(byteValue, &conf)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Dbuser, conf.Dbpass, conf.Dbhost, conf.Dbport, conf.Dbname)

	DB, err = sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	Secret = conf.Secret
	DevMode = conf.Sslsup

	if DevMode == "https" {
		IOURL = ""
		Secured = true
	} else {
		IOURL = ""
		Secured = false
	}

	Auth = ent.NetGsmAuth{
		User:   "user",
		Pass:   "pass",
		Header: "header",
		Num:    "num",
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
