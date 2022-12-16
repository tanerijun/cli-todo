# cli-todo

A simple CLI TODO list app

# Install

```
go install https://github.com/tanerijun/cli-todo
```

Or you can also clone the project and build it yourself. Doing this also allows you to change the binary name.

```
git clone https://github.com/tanerijun/cli-todo.git
cd cli-todo
go build -i -o custom_name
```

# Usage

1. View todos
   ```
   cli-todo ls
   ```
2. Add todo
   ```
   cli-todo add "Clean the ceiling"
   ```
3. Remove todo
   ```
   cli-todo rm 1
   ```

Use the `help` command for more details.
