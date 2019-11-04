package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/peterwade153/bucketlist/models"
)

// CreateItem Creates new itmes 
func CreateItem(w http.ResponseWriter, r *http.Request){

	var newItem models.Item

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	er := json.Unmarshal(reqBody, &newItem)
	if er != nil{
		w.WriteHeader(400)
		return
	}
	// update Items to add new item
	models.Items = append(models.Items, newItem)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(newItem)

}