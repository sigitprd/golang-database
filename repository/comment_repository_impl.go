package repository

import (
	"context"
	"database/sql"
	"golang-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	stmt, err := repo.DB.PrepareContext(ctx, query)

	if err != nil {
		return entity.Comment{}, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, comment.Email, comment.Comment)

	if err != nil {
		return entity.Comment{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entity.Comment{}, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ?"
	stmt, err := repo.DB.PrepareContext(ctx, query)

	if err != nil {
		return entity.Comment{}, err
	}

	defer stmt.Close()

	var comment entity.Comment

	err = stmt.QueryRowContext(ctx, id).Scan(&comment.Id, &comment.Email, &comment.Comment)

	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	stmt, err := repo.DB.PrepareContext(ctx, query)

	if err != nil {
		return []entity.Comment{}, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)

	if err != nil {
		return []entity.Comment{}, err
	}

	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		var comment entity.Comment
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			return []entity.Comment{}, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
