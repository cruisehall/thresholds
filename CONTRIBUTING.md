## Project Structure

```
thresholds/
├── cmd/            # Entrypoint applications
│   └── thresholds/ # Main application
│       ├── main.go
│       └── gen_signals_doc.go
├── internal/       # Private application code
├── pkg/            # Public reusable packages
├── go.mod
├── .gitignore
└── README.md
```

- Place your main application code in `cmd/thresholds/main.go`.
- Use `internal/` for private packages.
- Use `pkg/` for reusable libraries.

## Generating `signals/all.md`

Before running the documentation generator for the first time, initialize Go modules and install dependencies:

```sh
go mod init thresholds
go get gopkg.in/yaml.v3
```

Then generate the documentation:

```sh
go run cmd/thresholds/gen_signals_doc.go
```

This script reads all `*.signal.yaml` files in the `signals` directory and uses the `signal.tmpl` template to generate the documentation.