# docker-clean-tool

A CLI tool written in Go to list and delete all Docker images on your machine.

## What it does

- Connects to your local Docker daemon
- Lists all Docker images with their name and ID
- Deletes every image, showing a green ✓ for success and red ✗ for failures
- Prints a summary of how many images were removed

## Requirements

- Go 1.24+
- Docker or Rancher Desktop running locally

## Install

Clone the repo and install the binary:

```bash
git clone https://github.com/ev183/docker-clean-tool.git
cd docker-clean-tool
go install .
```

Make sure `~/go/bin` is on your PATH:

```bash
export PATH=$PATH:~/go/bin
```

Add that line to your `~/.zshrc` or `~/.bashrc` to make it permanent.

## Usage

```bash
docker-clean-tool
```

That's it. It will list all images then delete them all immediately.

## Example output

```
docker-clean-tool starting...
nginx:latest        a1b2c3d4e5f6
postgres:15         b2c3d4e5f6a7
<none>              c3d4e5f6a7b8

  ✓ Deleted: nginx:latest
  ✓ Deleted: postgres:15
  ✓ Deleted: <none>

Done! Removed: 3
```

## Built with

- [moby/moby](https://github.com/moby/moby) — Docker SDK for Go
- [fatih/color](https://github.com/fatih/color) — terminal color output