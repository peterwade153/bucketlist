package main

import (
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/peterwade153/bucketlist/controllers"
)

func TestCreateItem(t *testing.T){

	var dummydata = []byte (`{"id":5, "title":"title", "description":"description"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(dummydata))
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateItem)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusCreated{
		t.Errorf("wrong status code: got %v instead of %v", status, http.StatusCreated)
	}

}

func TestEditItem(t *testing.T){
	// creating item to edit
	var dummydata = []byte (`{"id":5, "title":"title", "description":"description"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(dummydata))
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateItem)
	handler.ServeHTTP(res, req)

	// editting the item
	var editdata = []byte (`{"id":5, "title":"title", "description":"description"}`)
	editreq, editerr := http.NewRequest("PUT", "/items/5", bytes.NewBuffer(editdata))
	if editerr != nil{
		t.Fatal(editerr)
	}
	editreq.Header.Set("Content-Type", "application/json")
	editres := httptest.NewRecorder()
	edithandler := http.HandlerFunc(controllers.EditItem)
	edithandler.ServeHTTP(editres, editreq)
	if status := editres.Code; status != http.StatusOK{
		t.Errorf("wrong status code: got %v instead of %v", status, http.StatusOK)
	}
}

func TestDeleteItem(t *testing.T){
	// creating item to edit
	var dummydata = []byte (`{"id":5, "title":"title", "description":"description"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(dummydata))
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateItem)
	handler.ServeHTTP(res, req)

	// deleting item
	delreq, delerr := http.NewRequest("DELETE", "/items/5", nil)
	if delerr != nil{
		t.Fatal(delerr)
	}

	delres := httptest.NewRecorder()
	delhandler := http.HandlerFunc(controllers.DeleteItem)
	delhandler.ServeHTTP(delres, delreq)
	if status := delres.Code; status != http.StatusOK{
		t.Errorf("wrong status code: got %v instead of %v", status, http.StatusOK)
	}
}

func TestGetallItems(t *testing.T){
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetallItems)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK{
		t.Errorf("wrong status code: got %v instead of %v", status, http.StatusOK)
	}
}

func TestGetItem(t *testing.T){
	req, err := http.NewRequest("GET", "/items/1", nil)
	if err != nil{
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetItem)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK{
		t.Errorf("wrong status code: got %v instead of %v", status, http.StatusOK)
	}
}
