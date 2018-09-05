package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template
var router *mux.Router

func init() {
	tpl = template.Must(template.ParseGlob("Views/*"))
}

func main() {
	addRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func addRoutes() {
	router = mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/jsontest", jsonTest).Methods("GET", "OPTIONS")
}

//Home loads React App
func Home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Home.html", nil)
}

func jsonTest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	json.NewEncoder(w).Encode("OK")
}
