package models

import (
	"database/sql"
	"time"
)

type Work struct {
	ID      int
	Content string
	Added   time.Time
}

type WorkModel struct {
	Db *sql.DB
}

func (m *WorkModel) TableExists() bool {
	sqlString := `SELECT * FROM work LIMIT 1`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return false
	}
	return true
}

func (m *WorkModel) CreateTable() error {
	sqlString := `CREATE TABLE work (
					"id" INTEGER PRIMARY KEY AUTOINCREMENT, 
					"content" TEXT NOT NULL,
					"added" DATETIME NOT NULL)`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return err
	}
	return nil
}

func (m *WorkModel) Insert(work string, date time.Time) error {
	sqlSmt := `INSERT INTO work 
				(content, added) 
				VALUES (?, ?)`
	_, err := m.Db.Exec(sqlSmt, work, date)
	return err
}

func (m *WorkModel) Update(w Work) error {
	sqlSmt := `UPDATE work SET content = ?, added = ? WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, w.Content, w.Added, w.ID)
	return err
}

func (m *WorkModel) Delete(id int) error {
	sqlSmt := `DELETE FROM work WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, id)
	return err
}

func (m *WorkModel) List(date time.Time) ([]Work, error) {

	from := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	to := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, time.Local)

	sqlSmt := `SELECT * FROM work WHERE added BETWEEN ? AND ?`
	rows, err := m.Db.Query(sqlSmt, from, to)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var works []Work
	for rows.Next() {
		var w Work
		err := rows.Scan(&w.ID, &w.Content, &w.Added)
		if err != nil {
			return nil, err
		}
		works = append(works, w)
	}
	return works, nil
}

func (m *WorkModel) Get(id int) (Work, error) {

	sqlSmt := `SELECT id, content, added FROM work WHERE id = ?`

	row := m.Db.QueryRow(sqlSmt, id)

	var w Work
	err := row.Scan(&w.ID, &w.Content, &w.Added)
	if err != nil {
		return Work{}, err
	}

	return w, nil

}
