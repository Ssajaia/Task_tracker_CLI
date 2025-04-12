project URL: [https://github.com/Ssajaia/Task_tracker_CLI.git](https://roadmap.sh/projects/task-tracker)

# Task Tracker CLI

A simple command-line application written in Go for managing tasks. It allows users to add, update, delete, and list tasks. Tasks are stored in a local JSON file.

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List all tasks or filter by status
- All data is stored in a local `tasks.json` file

## Getting Started

### Prerequisites

- Go 1.18 or higher

### Installation

1. Clone the repository or download the source code.
2. Build the CLI:

```bash
go build -o task-cli main.go
The binary task-cli will be created in the current directory.

Usage
Add a new task
bash
Copy
Edit
./task-cli add "Buy groceries"
Update a task
bash
Copy
Edit
./task-cli update 1 "Buy groceries and cook dinner"
Delete a task
bash
Copy
Edit
./task-cli delete 1
Mark task as in progress
bash
Copy
Edit
./task-cli mark-in-progress 1
Mark task as done
bash
Copy
Edit
./task-cli mark-done 1
List all tasks
bash
Copy
Edit
./task-cli list
List tasks by status
bash
Copy
Edit
./task-cli list done
./task-cli list todo
./task-cli list in-progress
Task Structure
Each task in tasks.json has the following structure:

json
Copy
Edit
{
  "id": 1,
  "description": "Example task",
  "status": "todo",
  "createdAt": "2025-04-12T14:00:00Z",
  "updatedAt": "2025-04-12T14:00:00Z"
}
Notes
If tasks.json does not exist, it will be created automatically.

The app uses the current directory to store the tasks.json file.

No external dependencies or libraries are used.

License
This project is open source and available under the MIT License.

yaml
Copy
Edit

--- 

Let me know if you want this saved into an actual `README.md` file!
