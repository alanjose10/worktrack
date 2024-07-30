# WorkTrack

WorkTrack is a command-line interface (CLI) tool designed to help you keep track of your daily tasks, manage your to-do list, generate reports of your completed work, and summarize your activities for stand-ups, sprints, and yearly reviews.

## Tech Stack

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper?tab=readme-ov-file)

## To Do

- Use https://github.com/dustin/go-humanize to format dates, etc.
- Support json output format via --json flag
- Custom help output

## Project Structure

```
worktrack/
│
├── cmd/
│   └── worktrack/
│       └── main.go
├── internal/
│   ├── add/
│   │   └── add.go
│   ├── list/
│   │   └── list.go
│   ├── todo/
│   │   ├── add.go
│   │   ├── list.go
│   │   └── remove.go
│   ├── blocker/
│   │   ├── add.go
│   │   ├── list.go
│   │   └── remove.go
│   ├── report/
│   │   └── report.go
│   ├── sprint/
│   │   └── sprint.go
│   ├── standup/
│   │   └── standup.go
│   └── storage/
│       ├── storage.go
│       └── fileutils.go
├── pkg/
│   ├── workitem/
│   │   └── workitem.go
│   ├── todoitem/
│   │   └── todoitem.go
│   └── blocker/
│       └── blocker.go
├── .gitignore
├── go.mod
└── README.md
```

## To do

```
**************************************************
Adding work items:
worktrack add "Implemented user authentication feature"

**************************************************

Add work item with a specific date:
worktrack add "Fixed bug in login page" --date 2021-01-01

**************************************************

Add todo item:
worktrack todo add "Update README.md"

**************************************************

Listing todo items:
worktrack todo list

Output:

WorkTrack - List of Work Items

Date       | Status   | Description
-----------|----------|----------------------------------------------
2023-07-21 | Completed| Implemented user authentication feature
2023-07-22 | Completed| Fixed bug in payment processing
2023-07-23 | Completed| Wrote unit tests for user service
2023-07-24 | Completed| Deployed new release to production
2023-07-25 | Completed| Updated project documentation
2023-07-26 | Pending  | Start working on the notification system
2023-07-27 | Pending  | Refactor database schema for performance
2023-07-28 | Pending  | Review pull requests for UI improvements

Total: 8 items (5 completed, 3 pending)

**************************************************

Marking todo item as done:
worktrack todo done 1

When a todo item is marked as done, it will altomatically be added to the list of completed work items and then removed from the todo list.

**************************************************

Add a blocker

worktrack blocker add "Waiting for API documentation from team"

**************************************************

List blockers

worktrack blocker list

Output:

WorkTrack - List of Blockers

ID  | Description
----|----------------------------------------------
1   | Unable to connect to the database
2   | Dependency on external API not resolved
3   | Lack of access to staging environment

Total: 3 blockers

**************************************************

Remove blocker:

worktrack blocker remove 1

**************************************************

Generating a report of completed work:
worktrack report --days 7

Output:

Report for the last 7 days:
1. 2023-07-21: Implemented user authentication feature
2. 2023-07-22: Fixed bug in payment processing
3. 2023-07-23: Wrote unit tests for user service
4. 2023-07-24: Deployed new release to production
5. 2023-07-25: Updated project documentation

**************************************************

Stand-up summary:
worktrack standup

Output:

Stand-Up Summary:
Yesterday:
1. Implemented user authentication feature
2. Fixed bug in payment processing

Today:
1. Write unit tests for user service
2. Start working on the notification system

Blockers:
None

Todo:
None

**************************************************

Sprint summary:
worktrack sprint

Output:

Sprint Summary (2023-07-10 to 2023-07-24):
1. Implemented user authentication feature
2. Fixed bug in payment processing
3. Wrote unit tests for user service
4. Deployed new release to production
5. Updated project documentation
6. Completed initial setup for the notification system

**************************************************

Yearly summary:
worktrack year

**************************************************

```

## Cli Commands

```bash
worktrack [command]

Available Commands:
  add         Add a work item
  todo        Manage to-do items
  blocker     Manage blockers
  report      Generate a detailed report of completed work
  standup     Generate a summary for daily stand-ups
  sprint      Generate a summary for the last sprint
  year        Generate a summary for the current year
  help        Help about any command
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contact
For questions or suggestions, please open an issue on GitHub.