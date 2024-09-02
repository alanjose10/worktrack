package main

import (
	"fmt"
	"time"

	"github.com/alanjose10/worktrack/internal/helpers"
)

func (app *application) addTodo() error {

}

func (app *application) listItems(fromDate time.Time, toDate time.Time) string {

	return fmt.Sprintf("Listing items between %s and %s\n", helpers.GetHumanDate(fromDate), helpers.GetHumanDate(toDate))
}
