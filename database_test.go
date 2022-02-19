package golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	// make sql connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()

}
