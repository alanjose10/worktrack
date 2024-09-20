package main

import (
	"fmt"
	"os"
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

	var selectedTodoId string
	var confirm bool
	// var textInput string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
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
				Title("Please select the todo item to mark as done").
				Value(&selectedTodoId),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Value(&confirm).
				Title("Are you sure you want to mark this item as done?").
				Validate(func(b bool) error {
					if !b {
						os.Exit(0)
					}
					return nil
				}).
				Affirmative("Yes!").
				Negative("No."),
		),
		// huh.NewGroup(
		// 	huh.NewText().
		// 		Title("Do you want to add the completed todo to work items?").
		// 		Value(&textInput).
		// 		Placeholder(fmt.Sprintf("Completed: %s", selectedTodoId)).
		// 		PlaceholderFunc(func() string {
		// 			for _, todo := range todos {
		// 				if fmt.Sprintf("%d", todo.ID) == selectedTodoId {
		// 					return fmt.Sprintf("Completed: %s", todo.Content)
		// 				}

		// 			}
		// 			return ""
		// 		}, &selectedTodoId).
		// 		Validate(func(s string) error {
		// 			if s == "" {
		// 				return fmt.Errorf("please provide a valid text")
		// 			}
		// 			return nil
		// 		}),
		// ),
	)

	err = form.Run()

	if err != nil {
		return fmt.Errorf("something went wrong: %s", err)
	}

	if !confirm {
		return nil
	}

	selectedTodoIdInt, err := strconv.Atoi(selectedTodoId)
	if err != nil {
		return fmt.Errorf("error converting id to int: %w", err)
	}
	selectedTodo, err := app.todoModel.GetById(selectedTodoIdInt)
	if err != nil {
		return fmt.Errorf("error getting todo by id: %w", err)
	}

	selectedTodo.Done = true
	err = app.todoModel.Update(selectedTodo)
	if err != nil {
		return fmt.Errorf("error updating todo: %w", err)
	}

	// textInput = fmt.Sprintf("%s (added via todo completion)", textInput)
	// err = app.workModel.Insert(textInput, helpers.GetCurrentDate())
	// if err != nil {
	// 	return fmt.Errorf("error inserting work: %w", err)
	// }

	fmt.Fprintf(os.Stdout, "Todo item marked as done: %s\n", selectedTodo.Content)

	return nil
}

func (app *application) undoTodo(from time.Time, to time.Time) error {

	todos, err := app.todoModel.ListBetween(from, to)
	if err != nil {
		return fmt.Errorf("error listing todos: %w", err)
	}

	var selectedTodoId string
	var confirm bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(10).
				OptionsFunc(func() []huh.Option[string] {
					var opts []huh.Option[string]
					for _, todo := range todos {
						if todo.Done {
							opts = append(opts, huh.NewOption(todo.Content, fmt.Sprintf("%d", todo.ID)))
						}

					}
					return opts
				}, nil).
				Title("Please select the todo item to mark as undone").
				Value(&selectedTodoId),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Value(&confirm).
				Title("Are you sure you want to mark this item as undone?").
				Validate(func(b bool) error {
					if !b {
						os.Exit(0)
					}
					return nil
				}).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err = form.Run()

	if err != nil {
		return fmt.Errorf("something went wrong: %s", err)
	}

	if !confirm {
		return nil
	}

	selectedTodoIdInt, err := strconv.Atoi(selectedTodoId)
	if err != nil {
		return fmt.Errorf("error converting id to int: %w", err)
	}
	selectedTodo, err := app.todoModel.GetById(selectedTodoIdInt)
	if err != nil {
		return fmt.Errorf("error getting todo by id: %w", err)
	}

	selectedTodo.Done = false
	err = app.todoModel.Update(selectedTodo)
	if err != nil {
		return fmt.Errorf("error updating todo: %w", err)
	}

	fmt.Fprintf(os.Stdout, "Todo item marked as undone: %s\n", selectedTodo.Content)

	return nil
}
