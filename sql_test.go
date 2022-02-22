package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
			id, name  string
			email     sql.NullString
			balance   int32
			rating    float64
			birthDate sql.NullTime
			createdAt time.Time
			married   bool
		)
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==========================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
	}

	fmt.Println("Success query all customer")
}

func TestSqlInjection(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "joko' OR 1=1 --"
	// password := "joko' OR 1=1 --"
	username := "admin'; #"
	// password := "admin"

	// username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	// Execute the query
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success |", username)
	} else {
		fmt.Println("Authentication failed")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "joko' OR 1=1 --"
	// password := "joko' OR 1=1 --"
	// username := "admin'; #"

	username := "admin"
	password := "admin"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	// Execute the query
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success |", username)
	} else {
		fmt.Println("Authentication failed")
	}
}

func TestExecSqlParameter(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "sigit"
	password := "sigit"

	query := "INSERT INTO user (username, password) VALUES (?, ?)"

	// Execute the query
	result, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
	fmt.Println(result)
}

func TestAutoIncrement(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "joko@gmail.com"
	comment := "Test Comment"

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	// Execute the query
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id :", insertId)
}

func TestPrepareStatement(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	// Prepare statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "sigit" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Comment -" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", lastInsertId)
	}

}

func TestTransaction(t *testing.T) {
	// get sql connection
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	/* do transaction */

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	// Prepare statement
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "sigit" + strconv.Itoa(i) + "@gmail.com"
		comment := "Test Comment -" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", lastInsertId)
	}

	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
