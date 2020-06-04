package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.homedepot.com/EMC4JQ2/docker-go-api/database"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
)

var handlers products.Product

func setup() {
	conn := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USER"),
		"pw":       os.Getenv("DB_PASSWORD"),
		"database": os.Getenv("DB_DATABASE"),
	}
	//fmt.Printf("%+v", conn)
	caller := database.DBCalls{}
	database.Init(caller, conn)
	handlers = products.GetHandlers(caller)
}

/* Don't forget to `source .env` */

func TestGetAll(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(handlers.GetAll))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var arrOfProducts []products.Product
	if err := json.Unmarshal([]byte(string(actual)), &arrOfProducts); err != nil {
		fmt.Println("error was ", err.Error())
		return
	}

	//fmt.Printf("%+v", arrOfProducts)
	assert.Equal(t, 9, len(arrOfProducts))
	assert.Equal(t, "brand A", arrOfProducts[0].Brand)
	assert.Equal(t, "Description 1", arrOfProducts[0].Description)
	assert.Equal(t, "THD-1", arrOfProducts[0].Sku)
	assert.Equal(t, 999, arrOfProducts[0].Price)
}
func TestGetOne(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(handlers.GetOne))
	defer ts.Close()
	res, err := http.Get(ts.URL + "/api/products/1")
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	var prod products.Product
	err = json.Unmarshal([]byte(string(actual)), &prod)
	assert.Nil(t, err)

	assert.Equal(t, 1, prod.Id)
	assert.Equal(t, "Product 1", prod.Name)
	assert.Equal(t, "brand A", prod.Brand)
	assert.Equal(t, "Description 1", prod.Description)
	assert.Equal(t, "THD-1", prod.Sku)
	assert.Equal(t, 4.2, prod.Rating)
	assert.Equal(t, 999, prod.Price)
}

func TestPostNew(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(handlers.PostNew))
	defer ts.Close()
	buf := bytes.NewBuffer([]byte(`{
		"name": "Power Driver",
		"brand": "DeWALT",
		"description": "MAX Lithium-Ion Brushless Cordless Compact 1/2 in. Drill Driver",
		"sku": "THD-10", 
		"rating": 4.8,
		"price": 9900,
		"department_id": 1
	}`))

	res, err := http.Post(ts.URL+"/api/products", "application/json", buf) // buf is already a pointer to a buffer.
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Contains(t, string(actual), "Created new item with id: ")
}

// Test update,
func TestUpdate(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(handlers.Update))
	buf := bytes.NewBuffer([]byte(`{
		"name": "Power Driver",
		"brand": "DeWALT",
		"description": "MAX Lithium-Ion Brushless Cordless Compact 1/2 in. Drill Driver",
		"sku": "THD-10", 
		"rating": 4.8,
		"price": 9900,
		"department_id": 1
	}`))
	client := &http.Client{}
	req, err := http.NewRequest("PUT", ts.URL+"/api/products/1", buf)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Contains(t, string(actual), "Updated item with id: 1")
}

func TestDelete(t *testing.T) {
	setup()
	ts := httptest.NewServer(http.HandlerFunc(handlers.Remove))
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", ts.URL+"/api/products/9", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Contains(t, string(actual), "Deleted item with id: 9")
}

