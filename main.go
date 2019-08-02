package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var router *mux.Router

func init() {}

func main() {
	addRoutes()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func addRoutes() {
	router = mux.NewRouter()
	router.HandleFunc("/uploadImage", uploadImage)
	router.HandleFunc("/getImagesForHome", getImagesForHome).Methods("GET")
}

func getImagesForHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit")
	// fakes call

	//var path = "Images/test-media"

	files, err := ioutil.ReadDir("./Images/test-media")
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(files)
}

// initImages fakes getting the images
func initImages() {}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	//interrogate r
	fmt.Println("uploadImage")

	json.NewEncoder(w).Encode("OK")
}
