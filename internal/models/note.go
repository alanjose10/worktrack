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
