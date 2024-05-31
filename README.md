# c.l.i.g.o

c.l.i.g.o is a command-line interface (CLI) application built in Go (Golang) to manage task lists. With c.l.i.g.o, you can easily add, delete, list, and complete tasks directly from your terminal.

## Features

- **Add Task**: Add a new task to your task list.
- **Delete Task**: Remove an existing task from your task list.
- **List Tasks**: Display all tasks in your task list.
- **Complete Task**: Mark a task as complete.

## Installation

To install c.l.i.g.o, you need to have Go installed on your machine. If you don't have Go installed, you can download it from [here](https://golang.org/dl/).

1. Clone the repository:
    ```sh
    git clone https://github.com/FilledEther20/c.l.i.g.o.git
    ```

2. Navigate to the project directory:
    ```sh
    cd c.l.i.g.o
    ```

3. Build the application:
    ```sh
    go build -o c.l.i.g.o
    ```

## Usage

### Add Task

To add a task, use the `add` command followed by the task description:
```sh
go run main.go -add your_task
```
### Delete Task

To delete a task, use the -delete flag followed by the task ID:
```sh
go run main.go -delete=1
```
### List Tasks

To list all tasks, use the -list flag:
```sh
go run main.go -list
```
### Complete Task

To mark a task as complete, use the -complete flag followed by the task ID:
```sh
go run main.go -complete=1
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

