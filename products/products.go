package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Id           int
	Name         string
	Brand        string
	Description  string
	Sku          string
	Rating       float64
	Price        int
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	DepartmentId int       `json:"department_id" db:"department_id"`
}

var allproducts = []Product{
	Product{0, "Cordless Jig Saw", "Milwaukee", "M18 FUEL 18-Volt Lithium-Ion Brushless Cordless Jig Saw (Tool Only)", "2737-20", 5.0, 199, time.Now(), time.Now(), 1},
	Product{1, "SAWZALL Saw Blades", "SAWZALL", "Demolition Nail-Embedded Wood and Metal Cutting Bi-Metal Reciprocating Saw Blade Set", "49-22-5670", 4.5, 202, time.Now(), time.Now(), 1},
}
var maxId = 1

func (p *Product) GetAll(w http.ResponseWriter, r *http.Request) {
	jsn, err := json.Marshal(allproducts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

func (p *Product) PostNew(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Failed to decode "+err.Error(), http.StatusBadRequest)
		return
	}
	maxId++
	p.Id = maxId
	allproducts = append(allproducts, *p)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Created new item with id: %v", p.Id)))
}

func (p *Product) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/products/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	found := false
	for _, prod := range allproducts {
		if prod.Id == id {
			p = &prod
			found = true
		}
	}
	if !found {
		http.Error(w, "product with that id not found.", http.StatusBadRequest)
		return
	}
	jsn, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

func (p *Product) Update(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Failed to decode "+err.Error(), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/products/"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	found := false
	for i, prod := range allproducts {
		if prod.Id == id {
			allproducts[i] = *p
			found = true
		}
	}
	if !found {
		http.Error(w, "product with that id not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Updated item with id: %v", p.Id)))
}

func (p *Product) Remove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/products/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	found := false
	newProducts := []Product{}
	for _, v := range allproducts {
		if v.Id != id {
			newProducts = append(newProducts, v)
		} else {
			found = true
		}
	}
	if !found {
		http.Error(w, "Product with that id does not exist", http.StatusBadRequest)
		return
	} else {
		allproducts = newProducts
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted item with id: %v", p.Id)))
}
