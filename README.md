# TODO App CLI (Golang)
![Screenshot](image.png)
This is a command-line interface (CLI) application built with Go (Golang) that allows you to manage your TODO list. You can add, view, update, and delete tasks from the command line.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Contributing](#contributing)

## Installation

To use this TODO App CLI, make sure you have Go installed on your machine. Then, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/amirhasanzadehpy/todo-app-cli.git
```

2. Navigate to the project directory:

```bash
cd todo-app-cli
```

3. Build the application:

```bash
go build ./cmd/todo/
```

4. Run the application:

```bash
./todo
```
OR

Move the generated binary to a directory in your system's PATH to make it globally accessible:

```bash
sudo mv todo /usr/local/bin
```
Note: You might need to use sudo to move the binary to a system-wide directory.

Verify the installation by running:

```bash
todo -v
```

## Usage

Once you have installed and run the application, you can start managing your TODO list using various commands.

To view all available commands and their usage, run:

```bash
todo -help
```

To add a new task:
```bash
todo -add Buy groceries
```

To list all tasks:
```bash
todo -list
```

To mark a task as completed:
```bash
todo -complete=1
```

To delete a task:
```bash
todo -delete=2
```

## Commands

The following are the available commands provided by the TODO App CLI:

- `add` - Add a new task to the TODO list.
- `list` - View all tasks in the TODO list.
- `complete` - Complete a task from the TODO list.
- `delete` - Delete a task from the TODO list.
- `help` - Display help information about the CLI and its commands.

## Contributing

Contributions are welcome! If you find any issues with the TODO App CLI or would like to suggest improvements, please open an issue or submit a pull request on the GitHub repository.
