package narad

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

type NaradDB struct {
	db_ref     *sql.DB
	data_store map[string]string
}

type db_config struct {
	User_name string
	Password  string
	Database  string
}

type config struct {
	Db db_config
}

func getDBCredentials() config {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print("Error:", err)
	}

	var madras_config config
	err = json.Unmarshal(content, &madras_config)
	if err != nil {
		fmt.Print("Error:", err)
	}
	return madras_config
}

func Connect() (*NaradDB, error) {
	conf := getDBCredentials()
	params := fmt.Sprintf("user=%s dbname=%s  password=%s",
		conf.Db.User_name, conf.Db.Database, conf.Db.Password)

	db, err := sql.Open("postgres", params)
	db_struct := new(NaradDB)
	db_struct.db_ref = db
	return db_struct, err
}

func (n *NaradDB) Find(short_code string) (string, bool) {
	return "http://manyu.in", true
}

func (n *NaradDB) Register(url string) (string, bool) {
	return "prgNw0D", false
}
