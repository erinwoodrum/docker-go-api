/* Unit Tests*/
package products_test

import (
	"database/sql"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
)

type mockCaller struct{}

/*
AddToDB(string)(string, error)
	GetFromDB(string)(*sql.Rows, error)
	GetOneFromDB(string)*sql.Row
	AlterInDB(string) error
*/
var Query string
var called int
var toErr bool

func (mc mockCaller) AddToDB(query string) (string, error) {
	Query = query
	called += 1
	if toErr {
		return "", errors.New("some error")
	}
	return "1234", nil
}

func (mc mockCaller) GetFromDB(query string) (*sql.Rows, error) {
	Query = query
	called += 1
	if toErr {
		return nil, errors.New("some error")
	}
	return nil, nil
}

func (mc mockCaller) GetOneFromDB(query string) *sql.Row {
	Query = query
	called += 1
	return nil
}

func (mc mockCaller) AlterInDB(query string) error {
	Query = query
	called += 1
	if toErr {
		return errors.New("some error")
	}
	return nil
}

func TestGetAll(t *testing.T) {
	p := products.GetHandlers(mockCaller{})

	/*
	  Expect it to call GetFromDB,
	  call it once,
	  and check the query string.
	  It should return 200 code.
	  and value a.
	*/

	req := httptest.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	called = 0
	p.GetAll(w, req)
	assert.Equal(t, 1, called)
	assert.Equal(t, "select * from products;", Query)
	assert.Equal(t, 200, w.Code)

	// Now let's force an error.
	// We expect an error code to get returned.

	req = httptest.NewRequest("GET", "/api/products", nil)
	w = httptest.NewRecorder()
	called = 0
	toErr = true
	p.GetAll(w, req)
	assert.Equal(t, 1, called)
	assert.Equal(t, 500, w.Code)
}

// Rest of unit tests would follow here. 

