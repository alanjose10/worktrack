package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/alanjose10/worktrack/internal/models"
	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

type application struct {
	dataPath  string
	taskModel *models.TaskModel
}

func setupPath() string {
	// get XDG paths
	scope := gap.NewScope(gap.User, "worktrack")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}

	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		taskDir, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

	}

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
	db, err := sql.Open("sqlite3", filepath.Join(path, "tasks.db"))
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
	if !app.taskModel.TableExists() {
		if err := app.taskModel.CreateTable(); err != nil {
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

	app := &application{
		dataPath:  path,
		taskModel: &models.TaskModel{Db: db},
	}
	if err := app.initializeDB(); err != nil {
		log.Fatal(err)
	}

	buildRootCommand(app).Execute()

}
