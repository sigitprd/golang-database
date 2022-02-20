package golang_database

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer (id, name) VALUES ('joko', 'JOKO')"

	// Execute the query
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
	fmt.Println(result)
}

func TestQuerySql(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM customer"

	// Execute the query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Success query all customer")
}
