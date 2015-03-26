package main

import (
	"fmt"
	"net/http"
	"strings"
)

var data_store map[string]string

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	paths := strings.Split(string(r.URL.Path[1:]), "/")
	key := paths[1]
	value := paths[2]
	val, ok := data_store[key]
	if ok {
		fmt.Fprintf(w, "Key - %s already present and set to %s", key, val)
	} else {
		data_store[key] = value
		fmt.Fprintf(w, "Key - %s not present and now set to %s", key, value)
	}
}
func retrivalHandler(w http.ResponseWriter, r *http.Request) {
	key := string(r.URL.Path[1:])
	val, ok := data_store[key]
	if ok {
		fmt.Fprint(w, val)
	} else {
		fmt.Fprintf(w, "Key %s - NOT PRESENT", key)
	}
}

func main() {
	data_store = make(map[string]string)
	http.HandleFunc("/register/", registrationHandler)
	http.HandleFunc("/", retrivalHandler)
	http.ListenAndServe(":8080", nil)
}
