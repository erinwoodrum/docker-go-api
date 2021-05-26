package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	
	"testing"

	"github.com/stretchr/testify/assert"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
)

var handlers products.Product

func setup() {
	handlers = products.Product{}
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
	assert.Equal(t, 2, len(arrOfProducts))
	assert.Equal(t, "Milwaukee", arrOfProducts[0].Brand)
	assert.Equal(t, "M18 FUEL 18-Volt Lithium-Ion Brushless Cordless Jig Saw (Tool Only)", arrOfProducts[0].Description)
	assert.Equal(t, "2737-20", arrOfProducts[0].Sku)
	assert.Equal(t, 199, arrOfProducts[0].Price)
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
	assert.Equal(t, "SAWZALL Saw Blades", prod.Name)
	assert.Equal(t, "SAWZALL", prod.Brand)
	assert.Equal(t, "Demolition Nail-Embedded Wood and Metal Cutting Bi-Metal Reciprocating Saw Blade Set", prod.Description)
	assert.Equal(t, "49-22-5670", prod.Sku)
	assert.Equal(t, 4.5, prod.Rating)
	assert.Equal(t, 202, prod.Price)
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
	req, err := http.NewRequest("DELETE", ts.URL+"/api/products/1", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)

	actual, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Contains(t, string(actual), "Deleted item with id: 1")
}

