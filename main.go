//todo
// figure out how to get routes to be case insensitive, double fns?
// think about a config file
// json encoder needs to camelcase things

package main

import (
	"dimgur-go/models"
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

var homeImages []models.ImageModel

func init() {
	homeImages = append(homeImages, models.ImageModel{URL: "/static/media/image1.8e1a0adb.png", Index: 0})
}

func main() {
	addRoutes()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func addRoutes() {
	router = mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/uploadImage", uploadImage).Methods("POST", "OPTIONS")
	router.HandleFunc("/getImages", getImages)
}

//Home loads React App
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("home hit")

}

func getImages(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getImages hit")

	json.NewEncoder(w).Encode(homeImages)
}

// initImages fakes getting the images
func initImages() {}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	//interrogate r
	fmt.Println("uploadImage")

	json.NewEncoder(w).Encode("OK")
}
