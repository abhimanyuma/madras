package main

import (
	"fmt"
	"github.com/abhimanyuma/madras/narad"
	"log"
	"net/http"
	"strings"
)

var data_store *narad.NaradDB
var db_error error

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(string(r.URL.Path[1:]), "/")
	url := paths[1]
	val, present := data_store.Register(url)

	if present {
		fmt.Fprintf(w, "Url - %s already present and set to %s", url, val)
	} else {
		fmt.Fprintf(w, "Url - %s not present and now set to %s", url, val)
	}

}

func retrivalHandler(w http.ResponseWriter, r *http.Request) {
	key := string(r.URL.Path[1:])
	val, present := data_store.Find(key)

	if present {
		fmt.Fprint(w, val)
	} else {
		fmt.Fprintf(w, "Key %s - NOT PRESENT", key)
	}
}

func main() {
	data_store, db_error = narad.Connect()

	if db_error != nil {
		log.Fatal(db_error)
		return
	}

	http.HandleFunc("/register/", registrationHandler)
	http.HandleFunc("/", retrivalHandler)
	http.ListenAndServe(":8080", nil)
}
