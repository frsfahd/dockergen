*This is a submission for the [GitHub Copilot CLI Challenge](https://dev.to/challenges/github-2026-01-21)*

## What I Built
Dockergen, a CLI tool that generates Dockerfile/Containerfile templates in seconds. It started as a way to remove repetitive setup work for Node.js and Python projects. The MVP focuses on speed and correctness: an interactive generator and an auto-detection mode that inspects a project to pick sensible defaults.

## Demo
Repo: https://github.com/frsfahd/dockergen



Quick demo:
```
dockergen generate
dockergen autogen
```
Both commands output a production-ready multi-stage Dockerfile with defaults tuned to the detected runtime.

`generate` is interactive: it asks you questions and uses your answers (or flags) to build the Dockerfile.

`autogen` is automatic: it inspects the current project (package.json, requirements, lockfiles, version files) to infer language/framework, ports, start/build commands, and then generates the Dockerfile with those defaults.

![dockergen generate](https://raw.githubusercontent.com/frsfahd/dockergen/main/screenshots/generate-command.png)
![dockergen autogen](https://raw.githubusercontent.com/frsfahd/dockergen/main/screenshots/autogen-command.png)

## My Experience with GitHub Copilot CLI
Copilot CLI helped me iterate fast. I used it to scaffold the Cobra-based commands, design the generator template flow, and build a lightweight auto-detection module. It was especially helpful for wiring the CLI flags and keeping the implementation focused on MVP scope while still following best practices.

<!-- Don't forget to add a cover image (if you want). -->


<!-- Thanks for participating! -->