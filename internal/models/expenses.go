package models

import (
	"database/sql"
	"time"
)

type ExpenseModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Expense, error)
	Latest() ([]*Expense, error)
}

type Expense struct {
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

func (m *ExpenseModel) Latest() ([]*Expense, error) {
	stmt := "SELECT id, category_id, title, amount, note, spent_on, created FROM expenses WHERE user_id = ? ORDER BY spent_on DESC, id DESC LIMIT 10"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expenses := []*Expense{}

	for rows.Next() {
		e := &Expense{}
		err = rows.Scan(&e.ID, &e.CategoryID, &e.Title,
			&e.Amount, &e.Note, &e.SpentOn, &e.Created)

		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (m *ExpenseModel) Insert(userID int, categoryID int, title string,
	amount int64, note string, spentOn time.Time) (int, error) {

	stmt := "INSERT INTO expenses (user_id, category_id, title, amount, note, spent_on, created) VALUES (?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())"

	var nullNote sql.NullString
	if note != "" {
		nullNote = sql.NullString{String: note, Valid: true}
	}

	result, err := m.DB.Exec(stmt, userID, categoryID, title, amount, nullNote, spentOn)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
