// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package generated

import (
	"database/sql"
)

type Comment struct {
	ID     uint64
	Body   string
	PostID sql.NullInt32
	UserID sql.NullInt32
}

type Post struct {
	ID     uint64
	Title  string
	Body   string
	UserID sql.NullInt32
}

type User struct {
	ID       uint64
	Username string
	Password string
}
