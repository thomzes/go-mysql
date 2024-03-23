package repository

import (
	"context"
	"fmt"
	gomysql "go-mysql"
	"go-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(gomysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment {
		Email : "test@gmail.com",
		Comment: "Test Comment",
	}

	result, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	CommentRepository := NewCommentRepository(gomysql.GetConnection())

	comment, err := CommentRepository.FindById(context.Background(), 33)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	CommentRepository := NewCommentRepository(gomysql.GetConnection())

	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}

}