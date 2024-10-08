package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/alanjose10/worktrack/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	version      string
	dataPath     string
	workModel    *models.WorkModel
	todoModel    *models.TodoModel
	blockerModel *models.BlockerModel
}

func setupPath() string {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	taskDir := filepath.Join(userHomeDir, ".worktrack")
	if _, err := os.Stat(taskDir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(taskDir, 0o770)
		} else {
			log.Fatal(err)
		}
	}

	return taskDir
}

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath.Join(path, "data.db"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (app *application) initializeDB() error {

	if !app.workModel.TableExists() {
		if err := app.workModel.CreateTable(); err != nil {
			return err
		}
	}

	if !app.todoModel.TableExists() {
		if err := app.todoModel.CreateTable(); err != nil {
			return err
		}
	}

	if !app.blockerModel.TableExists() {
		if err := app.blockerModel.CreateTable(); err != nil {
			return err
		}
	}

	return nil
}

func main() {

	path := setupPath()

	db, err := openDB(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	version := "1.3.0" // Define version as a constant at package level

	app := &application{
		version:      version,
		dataPath:     path,
		workModel:    &models.WorkModel{Db: db},
		todoModel:    &models.TodoModel{Db: db},
		blockerModel: &models.BlockerModel{Db: db},
	}
	if err := app.initializeDB(); err != nil {
		log.Fatal(err)
	}

	buildRootCommand(app).Execute()

}
