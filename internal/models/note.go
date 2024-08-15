package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID      int
	TaskID  int
	Content string
	AddedOn time.Time
}

type NoteModel struct {
	Db *sql.DB
}

func (nm *NoteModel) TableExists() bool {
	sqlString := `SELECT * FROM notes LIMIT 1`
	if _, err := nm.Db.Exec(sqlString); err != nil {
		return false
	}
	return true
}

func (nm *NoteModel) CreateTable() error {
	sqlString := `CREATE TABLE notes (
					"id" INTEGER PRIMARY KEY AUTOINCREMENT, 
					"task_id" INTEGER NOT NULL, 
					"content" TEXT NOT NULL,
					"added" DATETIME NOT NULL,
					FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
				)`
	if _, err := nm.Db.Exec(sqlString); err != nil {
		return err
	}
	return nil
}

func (nm *NoteModel) Insert(taskID int, content string) error {
	sqlString := `INSERT INTO notes (task_id, content, added) VALUES (?, ?, ?)`
	_, err := nm.Db.Exec(sqlString, taskID, content, time.Now())
	return err
}

func (nm *NoteModel) ListByTaskId(taskID int) ([]Note, error) {
	sqlString := `SELECT * FROM notes WHERE task_id = ?`
	rows, err := nm.Db.Query(sqlString, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		note := Note{}
		if err := rows.Scan(&note.ID, &note.TaskID, &note.Content, &note.AddedOn); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}
