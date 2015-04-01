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
	db_struct.data_store = make(map[string]string)
	return db_struct, err
}

func (n *NaradDB) Find(short_code string) (string, bool) {

	var url string
	val, present := n.data_store[short_code]

	if present {
		return val, true
	}

	query_string := fmt.Sprintf("SELECT url FROM redirects WHERE short_code like '%s'", short_code)
	err := n.db_ref.QueryRow(query_string).Scan(&url)

	switch {
	case err == sql.ErrNoRows:
		return "", false
	case err != nil:
		return err.Error(), false
	default:
		n.data_store[short_code] = url
		return url, true
	}
}

func (n *NaradDB) Register(url string) (string, bool) {
	return "prgNw0D", false
}
