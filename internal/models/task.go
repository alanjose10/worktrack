package models

import (
	"database/sql"
	"time"
)

type TaskModel struct {
	Db *sql.DB
}

func (tm *TaskModel) TableExists() bool {
	sqlString := `SELECT * FROM tasks LIMIT 1`
	if _, err := tm.Db.Exec(sqlString); err != nil {
		return false
	}
	return true
}

func (tm *TaskModel) CreateTable() error {
	sqlString := `CREATE TABLE tasks ("id" INTEGER PRIMARY KEY AUTOINCREMENT, 
					"name" TEXT NOT NULL, 
					"project" TEXT,
					"status" TEXT NOT NULL,
					"created" DATETIME )`
	if _, err := tm.Db.Exec(sqlString); err != nil {
		return err
	}
	return nil
}

type state int

const (
	todo state = iota
	inProgress
	done
)

// Returns the string representation of the status
func (s state) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

type Task struct {
	ID      int
	Name    string
	Project string
	Status  string
	Created time.Time
}

func (tm *TaskModel) Insert(name string, project string, status state) error {
	sqlSmt := `INSERT INTO tasks 
				(name, project, status, created) 
				VALUES (?, ?, ?, ?)`
	_, err := tm.Db.Exec(sqlSmt, name, project, status.String(), time.Now())
	return err
}

func (tm *TaskModel) Delete(id int) error {
	sqlSmt := `DELETE FROM tasks WHERE id = ?`
	_, err := tm.Db.Exec(sqlSmt, id)
	return err
}

func (m *TaskModel) List() ([]Task, error) {
	sqlSmt := `SELECT * FROM tasks`
	rows, err := m.Db.Query(sqlSmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (m *TaskModel) GetById(id int) (Task, error) {
	sqlSmt := `SELECT * FROM tasks WHERE id = ?`
	row := m.Db.QueryRow(sqlSmt, id)
	var t Task
	err := row.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
	if err != nil {
		return Task{}, err
	}
	return t, nil
}

func (m *TaskModel) GetByStatus(status state) ([]Task, error) {
	sqlSmt := `SELECT * FROM tasks WHERE status = ?`
	rows, err := m.Db.Query(sqlSmt, status.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.Project, &t.Status, &t.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
