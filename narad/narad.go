package narad

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type NaradDB struct {
	db_ref     *sql.DB
	data_store map[string]string
}

func Connect() (*NaradDB, error) {
	db, err := sql.Open("postgres", "user=madras dbname=madras password=Js6K9j9JajtVfkl6dsKoLia9Lf68dk805bKf")
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
