package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/peterwade153/bucketlist/controllers"
)

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "bucketlist home")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.EditItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", router))
}