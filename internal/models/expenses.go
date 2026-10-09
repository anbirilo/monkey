package models

import (
	"database/sql"
	"time"
)

type Expese struct {
	ID         int
	CategoryID int
	Title      string
	Amount     int64
	Note       sql.NullString
	SpentOn    time.Time
	Created    time.Time
}

type ExpenseModel struct {
	DB *sql.DB
}
