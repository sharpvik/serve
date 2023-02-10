# `serve`

Local static file server.

## Install

```bash
go install github.com/sharpvik/serve@latest
```

## Usage

```bash
NAME:
   serve - local static file server

USAGE:
   serve [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dir value, -d value   Serve files from the specified directory (default: ".")
   --port value, -p value  Bind server to custom port (default: 8888)
   --share, -s             Use host 0.0.0.0 instead of localhost (default: false)
   --help, -h              show help
```
