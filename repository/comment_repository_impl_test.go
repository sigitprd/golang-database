package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentRepository_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepository_FindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	id := int32(90)

	result, err := commentRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepository_FindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()

	results, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
