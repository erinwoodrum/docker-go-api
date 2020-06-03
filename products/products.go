package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	db "github.homedepot.com/EMC4JQ2/docker-go-api/database"
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



func (p *Product) GetAll(w http.ResponseWriter, r *http.Request) {
	var retArr []Product
	rows, err := db.GetFromDB("select * from products;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		one := Product{}
		if err := rows.Scan(&one.Id, &one.Name, &one.Brand, &one.Description, &one.Sku, &one.Rating, &one.Price, &one.CreatedAt, &one.UpdatedAt, &one.DepartmentId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		retArr = append(retArr, one)
	}

	jsn, err := json.Marshal(retArr)
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
	query := fmt.Sprintf("INSERT INTO products (name, brand, description, sku, rating, price, department_id) VALUES ('%v', '%v', '%v', '%v', %v, %v, %v) returning id;", p.Name, p.Brand, p.Description, p.Sku, p.Rating, p.Price, p.DepartmentId)
	fmt.Println(query)
	id, err := db.AddToDB(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Created new item with id: %v", id)))
}

func (p *Product) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.Id = id
	query := fmt.Sprintf("SELECT * FROM products WHERE id = %v", id)
	row := db.GetOneFromDB(query)
	err = row.Scan(&p.Id, &p.Name, &p.Brand, &p.Description, &p.Sku, &p.Rating, &p.Price, &p.CreatedAt, &p.UpdatedAt, &p.DepartmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	query := fmt.Sprintf("UPDATE products SET name = '%v', brand = '%v', description = '%v', sku = '%v', rating = %v, price = %v, department_id = %v  WHERE id = %v;", p.Name, p.Brand, p.Description, p.Sku, p.Rating, p.Price, p.DepartmentId, p.Id)
	fmt.Println(query)
	err := db.AlterInDB(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Updated item with id: %v", p.Id)))
}

func (p *Product) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.Id = id
	query := fmt.Sprintf("DELETE FROM products WHERE id = %v ;", p.Id)
	err = db.AlterInDB(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted item with id: %v", p.Id)))
}
