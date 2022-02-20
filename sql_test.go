package golang_database

import (
	"context"
	"fmt"
	"testing"
	"time"

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

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id, "| Name:", name)
	}

	fmt.Println("Success query all customer")
}

func TestQuerySqlComplex(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	// Execute the query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id, name, email      string
			balance              int32
			rating               float64
			birthDate, createdAt time.Time
			married              bool
		)
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==========================")
		fmt.Println("Id:", id, "| Name:", name)
		fmt.Println("Email:", email, "| Balance:")
		fmt.Println("Rating:", rating, "| Birth Date:", birthDate)
		fmt.Println("Maried:", married, "| Created At:", createdAt)
	}

	fmt.Println("Success query all customer")
}
