package main

var (
	addCmdExamples = `
Add a work item with inline input

	worktrack add "Worked on documenting the APIs"

Add a todo item with inline input

	worktrack add todo "Update the README"

Add a blocker item with inline input

	worktrack add blocker "Waiting for the API keys from ops team"


Add a work item using prompt

	worktrack add

Add a todo item

	worktrack add todo

Add a blocker item

	worktrack add blocker

Add an item for yesterday

	worktrack add -y

Add an item for a specific date

	worktrack add --date 20-10-1994
	`
	listCmdExamples = `
List work items for the last 7 days:
  worktrack list -d 7

List todo items for the last 2 weeks:
  worktrack list todo -w 2

List blocker items for the last 3 months:
  worktrack list blocker -m 3

List work items for the last year:
  worktrack list -y 1
    `
)
