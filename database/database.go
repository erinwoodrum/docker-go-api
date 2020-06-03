package database

// Get db password from Env.
import (
	"database/sql"
	"fmt"
	// The below import is to run the pq package in the background.
	_ "github.com/lib/pq"
)

var db *sql.DB

// Init creates the database connection
func Init() {
	host := "localhost"
	port := "5555"
	user := "postgres"
	pw := "secret" // TODO: get this from environment
	database := "products"
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pw, database)
	var err error
	db, err = sql.Open("postgres", connString)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open: %v", err))
	}
	err = db.Ping()
	if err == nil {
		fmt.Println("Ping successful")
		return
	}
	return
}

// AddToDB will add a row to a database and return the id of the
// new thing.
func AddToDB(query string) (id string, err error) {
	row := db.QueryRow(query)
	err = row.Scan(&id)
	return
}

// GetFromDB will retreive multiple rows from the database.
func GetFromDB(query string) (*sql.Rows, error) {
	return db.Query(query)
}

// GetOneFromDB will retreive one row from the database
func GetOneFromDB(query string) *sql.Row {
	return db.QueryRow(query)
}

// AlterInDB attempts to change one row in the database.
func AlterInDB(query string) error {
	row := db.QueryRow(query)
	var result string
	if err := row.Scan(&result); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil
		}
		return err
	}
	return nil
}