package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)

}

// EditItem enables editing an Item
func EditItem(w http.ResponseWriter, r *http.Request){

	id := mux.Vars(r)["id"]

	itemID, _ := strconv.Atoi(id)

	var updatedItem models.Item

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	er := json.Unmarshal(reqBody, &updatedItem)
	if er != nil{
		w.WriteHeader(400)
		return
	}

	for i, item := range(models.Items){
		if item.ID == itemID{
				item.Title = updatedItem.Title
				item.Description = updatedItem.Description

				models.Items = append(models.Items[:i], item)
				w.Header().Add("Content-Type", "application/json")
				json.NewEncoder(w).Encode(item)
		}
	}
}

// DeleteItem deletes items
func DeleteItem(w http.ResponseWriter, r *http.Request){

	id := mux.Vars(r)["id"]

	itemID, _ := strconv.Atoi(id)

	for i, item := range(models.Items){
		if item.ID == itemID{
			models.Items = append(models.Items[:i], models.Items[i+1:]...)
		}
	}
}

// GetallItems returns all items
func GetallItems(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Items)
}

// GetItem returns an items
func GetItem(w http.ResponseWriter, r *http.Request){

	id := mux.Vars(r)["id"]

	itemID, _ := strconv.Atoi(id)

	for _, item := range(models.Items){
		if item.ID == itemID{
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
		}
	}
}
