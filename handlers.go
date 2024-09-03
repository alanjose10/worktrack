package main

import (
	"fmt"
	"time"

	"github.com/alanjose10/worktrack/internal/output"
)

func (app *application) listWork(from time.Time, to time.Time) error {
	items, err := app.workModel.ListBetween(from, to)
	if err != nil {
		return fmt.Errorf("error listing work: %w", err)
	}

	fmt.Printf("%s\n", output.BuildListWorkOutput(from, to, items))

	return nil
}

func (app *application) listTodo(from time.Time, to time.Time) error {

	todos, err := app.todoModel.ListBetween(from, to)
	if err != nil {
		return fmt.Errorf("error listing todos: %w", err)
	}
	fmt.Printf("%s\n", output.BuildListTodoOutput(from, to, todos))

	return nil
}

func (app *application) listBlocker(from time.Time, to time.Time) error {

	blockers, err := app.blockerModel.ListBetween(from, to)
	if err != nil {
		return fmt.Errorf("error listing blockers: %w", err)
	}
	fmt.Printf("%s\n", output.BuildListBlockerOutput(from, to, blockers))
	return nil
}
