package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var router *mux.Router

func init() {
	tpl = template.Must(template.ParseGlob("Views/*"))
}

func main() {
	addRoutes()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func addRoutes() {
	router = mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/uploadImage", uploadImage)
}

//Home loads React App
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	tpl.ExecuteTemplate(w, "Home.html", nil)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	//interrogate r
	fmt.Println("uploadImage")

	json.NewEncoder(w).Encode("OK")
}
