
All data is stored locally and never leaves the user's device.

It will be stored in json format.

Location: `.worktrack` directory in the user's home

Directory structure:
```
.worktrack/
│
├── 2023/
│   ├── 07/
│   │   ├── 21/
│   │   │   ├── work_items.json
│   │   │   ├── todos.json
│   │   │   └── blockers.json
│   │   ├── 22/
│   │   │   ├── work_items.json
│   │   │   ├── todos.json
│   │   │   └── blockers.json
│   │   └── ...
│   └── ...
└── 2024/
    └── ...
```

Files:

- `work_items.json`: Contains all work items for the day
- `todos.json`: Contains all todos for the day
- `blockers.json`: Contains all blockers for the day

File contents:

- `work_items.json`:
  ```json
  [
    {
      "id": 1,
      "description": "Implemented user authentication feature"
    },
    {
      "id": 2,
      "description": "Fixed bug in payment processing"
    }
  ]
  ```
- `todos.json`:
  ```json
  [
    {
      "id": 1,
      "description": "Write unit tests for user service"
    },
    {
      "id": 2,
      "description": "Start working on the notification system"
    }
  ]
  ```
- `blockers.json`:
  ```json
  [
    {
      "id": 1,
      "description": "Waiting for approval on design mockups"
    }
  ]
  ```
