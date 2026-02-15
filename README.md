# Dockergen
Dockergen is a CLI tool for generating optimized Dockerfile/Containerfile templates for common runtimes.

## Usage
```
Usage:
    dockergen [command]

Available Commands:
    autogen     Auto-detect project and generate a Dockerfile/Containerfile
    generate    Interactively generate a Dockerfile/Containerfile
    help        Help about any command

Flags:
    -h, --help   help for dockergen

Use "dockergen [command] --help" for more information about a command.
```

### dockergen generate
```          
Interactively generate a Dockerfile/Containerfile

Usage:
    dockergen generate [flags]

Flags:
            --alpine                   Use Alpine-based images when supported (default true)
            --app string               Application name
            --build string             Build command (optional)
    -h, --help                     help for generate
            --lang string              Runtime language (node|python)
            --no-expose                Do not include EXPOSE instruction
            --node-version string      Node.js version (e.g. 20)
            --output string            Output file path (default: Dockerfile)
            --package-manager string   Node.js package manager (npm|yarn|pnpm)
            --port int                 Application port to expose
            --python-version string    Python version (e.g. 3.12)
            --requirements string      Python requirements file name (default "requirements.txt")
            --start string             Start command
            --workdir string           Work directory inside container (default "/app")
```

### dockergen autogen
```            
Auto-detect project and generate a Dockerfile/Containerfile

Usage:
    dockergen autogen [flags]

Flags:
            --alpine                   Use Alpine-based images when supported (default true)
            --app string               Application name
            --build string             Build command (optional)
    -h, --help                     help for autogen
            --no-expose                Do not include EXPOSE instruction
            --node-version string      Node.js version (e.g. 20)
            --output string            Output file path (default: Dockerfile)
            --package-manager string   Node.js package manager (npm|yarn|pnpm)
            --port int                 Application port to expose
            --python-version string    Python version (e.g. 3.12)
            --requirements string      Python requirements file name
            --start string             Start command
            --workdir string           Work directory inside container (default "/app")
```

## Install
```
go install github.com/frsfahd/dockergen@latest
```
Make sure your `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH`.

## Features
- [x] `generate` : generate dockerfile/containerfile via interactive user prompt 
- [x] `autogen` : generate dockerfile/containerfile automatically based on detected framework & language of the project

## Language & Framework Support
- [x] `Node.js` :
    - [x] `Next.js`
    - [x] `Vite`
    - [x] `Express.js`
    - [x] `Hapi`
    - [x] `Fastify`
- [x] `Python` :
    - [x] `Flask`
    - [x] `Fastapi`
    - [x] `Django`

see `internal/catalog/languages.go`
