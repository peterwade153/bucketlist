package tests

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