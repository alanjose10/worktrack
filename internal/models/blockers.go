package models

import (
	"database/sql"
	"time"
)

type Blocker struct {
	ID       int
	Content  string
	Added    time.Time
	Resolved bool
}

type BlockerModel struct {
	Db *sql.DB
}

func (m *BlockerModel) TableExists() bool {
	sqlString := `SELECT * FROM blocker LIMIT 1`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return false
	}
	return true
}

func (m *BlockerModel) CreateTable() error {
	sqlString := `CREATE TABLE blocker (
					"id" INTEGER PRIMARY KEY AUTOINCREMENT, 
					"content" TEXT NOT NULL,
					"resolved" BOOLEAN NOT NULL DEFAULT 0,
					"added" DATETIME NOT NULL)`
	if _, err := m.Db.Exec(sqlString); err != nil {
		return err
	}
	return nil
}

func (m *BlockerModel) Insert(blocker string, date time.Time) error {
	sqlSmt := `INSERT INTO blocker 
				(content, added) 
				VALUES (?, ?)`
	_, err := m.Db.Exec(sqlSmt, blocker, date)
	return err
}

func (m *BlockerModel) Update(b Blocker) error {
	sqlSmt := `UPDATE blocker SET content = ?, added = ?, resolved = ? WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, b.Content, b.Added, b.Resolved, b.ID)
	return err
}

func (m *BlockerModel) Delete(id int) error {
	sqlSmt := `DELETE FROM blocker WHERE id = ?`
	_, err := m.Db.Exec(sqlSmt, id)
	return err
}

// List returns all unresolved blockers
func (m *BlockerModel) List() ([]Blocker, error) {

	sqlSmt := `SELECT id, content, added, resolved FROM blocker WHERE resolved = 0`

	rows, err := m.Db.Query(sqlSmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blockers []Blocker
	for rows.Next() {
		var b Blocker
		err := rows.Scan(&b.ID, &b.Content, &b.Added, &b.Resolved)
		if err != nil {
			return nil, err
		}
		blockers = append(blockers, b)
	}
	return blockers, nil
}

// List blockers between two dates
func (m *BlockerModel) ListBetween(fromDate time.Time, toDate time.Time) ([]Blocker, error) {

	sqlSmt := `SELECT id, content, added, resolved FROM blocker WHERE added BETWEEN ? AND ?`
	rows, err := m.Db.Query(sqlSmt, fromDate, toDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blockers []Blocker
	for rows.Next() {
		var b Blocker
		err := rows.Scan(&b.ID, &b.Content, &b.Added, &b.Resolved)
		if err != nil {
			return nil, err
		}
		blockers = append(blockers, b)
	}
	return blockers, nil
}
