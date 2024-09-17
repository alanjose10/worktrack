package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alanjose10/worktrack/internal/output"
	"github.com/charmbracelet/huh"
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

func (app *application) doTodo(from time.Time, to time.Time) error {

	todos, err := app.todoModel.ListBetween(from, to)
	if err != nil {
		return fmt.Errorf("error listing todos: %w", err)
	}

	doneItems := []string{}
	huh.NewMultiSelect[string]().
		Height(10).
		OptionsFunc(func() []huh.Option[string] {
			var opts []huh.Option[string]
			for _, todo := range todos {
				if !todo.Done {
					opts = append(opts, huh.NewOption(todo.Content, fmt.Sprintf("%d", todo.ID)))
				}

			}
			return opts
		}, nil).
		Title("Please select an item to mark as done. (Use SPACE to select)").
		Value(&doneItems).
		Limit(1).
		Run()

	if len(doneItems) > 0 {

		idInt, err := strconv.Atoi(doneItems[0])
		if err != nil {
			return fmt.Errorf("error converting id to int: %w", err)
		}
		todo, err := app.todoModel.GetById(idInt)
		if err != nil {
			return fmt.Errorf("error getting todo by id: %w", err)
		}

		var confirm bool
		prompt := huh.NewConfirm().
			Description(todo.Content).
			Value(&confirm).
			TitleFunc(func() string {
				title := "Are you sure you want to mark this item as done?"
				return title
			}, nil).
			Affirmative("Yes!").
			Negative("No.")

		err = prompt.Run()
		if err != nil {
			return err
		}

		if !confirm {
			return nil
		}
		todo.Done = true
		err = app.todoModel.Update(todo)
		if err != nil {
			return fmt.Errorf("error updating todo: %w", err)
		}
	}

	return nil
}

func (app *application) undoTodo(from time.Time, to time.Time) error {

	return nil
}
