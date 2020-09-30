# Reverse Shell Generator

A tool to generate reverse shell code with provided host ip address and port. Useful for e.g. CTF and HackTheBox challenges.

## Install

```bash
go get -u github.com/mskri/reverse-shell-generator
```

## Usage

```bash
reverse-shell-generator --host 10.168.13.20 --port 4444
```

Listen to connections with e.g.

```bash
nc -vnlp 4444
```

## Development

Build and run with

```bash
go run .
```

you can also pass the flags to `go run .`

```bash
go run . --host 127.0.0.1 --port 4444
```

or build binary and run it manually

```bash
go build .
./reverse-shell-generator --host 127.0.0.1 --port 4444
```
