package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func addRoutes() {
	router = mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/uploadImage", uploadImage).Methods("POST", "OPTIONS")
}

//Home loads React App
func Home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	//interrogate r
	fmt.Println("cat")

	setHeaders(w, r)
	json.NewEncoder(w).Encode("OK")
}

func setHeaders(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
