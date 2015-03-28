package narad

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type NaradDB *sql.DB

func Connect() (NaradDB, error) {
	db, err := sql.Open("postgres", "user=madras dbname=madras password=Js6K9j9JajtVfkl6dsKoLia9Lf68dk805bKf")
	return db, err
}
