// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package repository

import (
	"database/sql"
)

type Profile struct {
	ID        string
	Username  string
	Firstname string
	Lastname  sql.NullString
	Avatar    sql.NullString
}
