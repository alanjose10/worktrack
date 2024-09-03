## Commands

### Work

```bash

# Adding a work item
worktrack add "Implemented user authentication feature"

# Listing work items
worktrack list
```

### Todo

```bash

# Adding a todo item
worktrack add todo "Update README.md"

# Listing todo items
worktrack list todo

# Marking a todo item as done
worktrack todo --do
# This command will list all the todo items which are marked as undone from the last 7 days (by default)


# Marking a todo item as undone
worktrack todo --undo
# This command will list all the todo items which are marked as done from the last 7 days (by default)

```