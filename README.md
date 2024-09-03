# WorkTrack

# Installation

```bash
go install github.com/alanjose10/worktrack@<tag-name>

Example:
go install github.com/alanjose10/worktrack@v1.2.0


```

# Introduction

WorkTrack is a command-line interface (CLI) tool designed to help you keep track of your daily tasks, manage your to-do list, generate reports of your completed work, and summarize your activities for stand-ups, sprints, and yearly reviews.

## Technologies

- [Go](https://golang.org/)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper?tab=readme-ov-file)
- [SQLite](https://www.sqlite.org/index.html)



## Cli Commands

```bash
$ go run . help
A CLI work tracker tool to make sprints & standups easier

Usage:
  worktrack [flags]
  worktrack [command]

Available Commands:
  add         Add items
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List items
  version     Print the version of the application
  where       Show where your tasks are stored

Flags:
  -h, --help   help for worktrack

Use "worktrack [command] --help" for more information about a command.
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contact
For questions or suggestions, please open an issue on GitHub.
