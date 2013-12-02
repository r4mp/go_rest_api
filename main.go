package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {

	log.Println("Starting server...")

	m := mux.NewRouter()

	// Get all lists.
	m.HandleFunc("/", GetAllLists).Methods("GET")

	// Make a new list.
	m.HandleFunc("/", PostList).Methods("POST")

	// Singe list operations.
	m.HandleFunc("/{key}/", GetList).Methods("GET")
/*	m.HandleFunc("/{key}/", PutList).Methods("PUT")
	m.HandleFunc("/{key}/", DeleteList).Methods("DELETE")

	// Everything else fails.
	m.HandleFunc("/{path:.*}", gorca.NotFoundFunc)
*/

	log.Println("Now listening on port 8080")

	//http.HandleFunc("/", handler)
	http.Handle("/", m)
	http.ListenAndServe(":8080", nil)
}

func PrintStuff(body map[string][]string) {
    log.Println(body)
}

func GetAllLists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Thanks for the %s!", r.Method)
}

func PostList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	body := make(map[string][]string)
	body["foo"] = r.PostForm["foo"]
	//go PrintStuff(body)
	// TODO: Was ist der Unterschied zu der Zeile darueber?
	PrintStuff(body)

	b, err := json.Marshal("aaaaa")

	if err == nil {
		//w.Write([]byte(b))
		w.Write(b)
	} else {
		w.Write([]byte("An error occured!\n"))
	}

	//w.Write([]byte("Thanks\n"))
}

func GetList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "r -  %s, w - %s", r, w)
	vars := mux.Vars(r)
	w.Write([]byte(vars["key"]))
}
