## Utility Commands

```bash
go work init

go work use ./module-name

go mod init github.com/username/projectname
```

Use `go work init` when you want to create a Go workspace. Workspaces are useful when you need to work with multiple modules within the same project or development environment.

The command creates a `go.work` file.

Use `go work use` to add a module to the current workspace.

For example:

```bash
go work use ./api
go work use ./shared
```

This updates the `go.work` file and allows the modules in the workspace to reference each other locally without needing to publish them first.

Use `go mod init` to initialize a new Go module. This command creates a `go.mod` file that defines the module path and manages its dependencies.

It is recommended to use a module path that matches the repository where the project will be hosted, for example:

```bash
go mod init github.com/username/projectname
```

This allows other Go projects to import your module using its repository path.
