package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
	"log"
	"os"
	"try-go-sqlc/generated"
)

type User struct {
	ID       uint64
	Username string
}

type Post struct {
	ID    uint64
	Title string
}

type Comment struct {
	ID   int64 // 何も設定しないと sql.NullInt64 が使われる
	Body string
}

func main() {
	db, _ := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	defer db.Close()

	queries := generated.New(db)
	ctx := context.Background()

	// パラメータなし
	result1, _ := queries.GetUsersPostsCount(ctx)
	log.Println(result1)

	// パラメータあり
	result2, _ := queries.GetUsersByIDs(ctx, []uint64{1, 2, 3})
	log.Println(result2)

	// パラメータあり、JOIN
	result3, _ := queries.GetUserPostWithComments(ctx, generated.GetUserPostWithCommentsParams{
		UserID: 1,
		PostID: 1,
	})
	type PostWithComments struct {
		Post     Post
		Comments []Comment
	}
	type UserPostWithComments struct {
		User User
		PostWithComments
	}
	postGroups := lo.GroupBy(result3, func(row generated.GetUserPostWithCommentsRow) uint64 {
		return uint64(row.PostID)
	})
	r := lo.FilterMap(result3, func(row generated.GetUserPostWithCommentsRow, _ int) (UserPostWithComments, bool) {
		postRows := postGroups[row.PostID]
		comments := lo.FilterMap(postRows, func(row generated.GetUserPostWithCommentsRow, _ int) (Comment, bool) {
			if !row.CommentID.Valid {
				return Comment{}, false
			}
			return Comment{
				ID:   row.CommentID.Int64,
				Body: row.Body.String,
			}, true
		})
		return UserPostWithComments{
			User: User{
				ID:       row.UserID,
				Username: row.Username,
			},
			PostWithComments: PostWithComments{
				Post: Post{
					ID:    row.PostID,
					Title: row.Title,
				},
				Comments: comments,
			},
		}, true
	})
	log.Println(r)
}
