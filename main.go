package main

import (
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
}

//Home loads React App
func Home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}
