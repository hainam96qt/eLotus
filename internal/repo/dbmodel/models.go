// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
)

type User struct {
	ID       int32          `json:"id"`
	UserName sql.NullString `json:"user_name"`
	Password sql.NullString `json:"password"`
}
