package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID      int
	Content string
	Done    bool
	Added   time.Time
}

type TodoModel struct {
	Db *sql.DB
}

func (m *TodoModel) TableExists() bool {
	sqlString := `SELECT * FROM todo LIMIT 1`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return false
	}
	return true
}

func (m *TodoModel) CreateTable() error {
	sqlString := `CREATE TABLE todo (
					"id" INTEGER PRIMARY KEY AUTOINCREMENT, 
					"content" TEXT NOT NULL,
					"done" BOOLEAN NOT NULL DEFAULT 0,
					"added" DATETIME NOT NULL)`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return err
	}
	return nil
}

func (m *TodoModel) Insert(todo string, date time.Time) error {
	sqlSmt := `INSERT INTO todo 
				(content, added) 
				VALUES (?, ?)`
	_, err := m.Db.Exec(sqlSmt, todo, date)
	return err
}

func (m *TodoModel) Update(t Todo) error {
	sqlSmt := `UPDATE todo SET content = ?, done = ? WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, t.Content, t.Done, t.Added, t.ID)
	return err
}

func (m *TodoModel) Delete(id int) error {
	sqlSmt := `DELETE FROM todo WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, id)
	return err
}

// List returns all todos that are not done
func (m *TodoModel) List() ([]Todo, error) {

	sqlSmt := `SELECT id, content, done, added FROM todo WHERE done = 0`

	rows, err := m.Db.Query(sqlSmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Content, &t.Done, &t.Added)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// List todos between two dates
func (m *TodoModel) ListBetween(fromDate time.Time, toDate time.Time) ([]Todo, error) {

	sqlSmt := `SELECT id, content, done, added FROM todo WHERE added BETWEEN ? AND ?`
	rows, err := m.Db.Query(sqlSmt, fromDate, toDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.ID, &t.Content, &t.Done, &t.Added)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}
