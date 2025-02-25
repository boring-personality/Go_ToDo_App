```markdown
# Todo CLI

Manage tasks efficiently by adding, deleting, and marking them complete.

## Usage

```sh
todo [command]
```

## Available Commands

- `add`         Add a task to the list
- `complete`    Mark a task as complete
- `completion`  Generate the autocompletion script for the specified shell
- `delete`      Delete a task
- `help`        Help about any command
- `list`        List the tasks

## Flags

- `-h`, `--help`   Help for the `todo` command

## Examples

### Add a Task

```sh
todo add "Buy groceries"
```

### List Tasks

```sh
todo list
```

#### List Output

The `list` command displays tasks in the following format:

```plaintext
ID   Task             Created
1    yay this works   18 minutes ago
```

### List All Tasks

Use the `-a` or `--all` flag to show all tasks, including completed and incomplete ones:

```sh
todo list -a
```

### Complete a Task

```sh
todo complete 1
```

### Delete a Task

```sh
todo delete 1
```

## Help

Use `todo [command] --help` for more information about a command.

---
