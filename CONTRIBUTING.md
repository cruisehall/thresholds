
## Generating `signals/all.md`

Before running the documentation generator for the first time, initialize Go modules and install dependencies:

```sh
go mod init thresholds
go get gopkg.in/yaml.v3
```

Then generate the documentation:

```sh
go run gen_signals_doc.go
```

This script reads all `*.signal.yaml` files in the `signals` directory and uses the `signal.tmpl` template to generate the documentation.