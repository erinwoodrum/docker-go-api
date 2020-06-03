/* Unit Tests*/
package product_test

import (
	"testing"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
	"net/http/httptest"
)

func TestGetAll(t *testing.T){
	p := products.Product{}
	req := httptest.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	p.GetAll(w, req)
	
}

