# RSG - Reverse Shell Generator

A tool to generate reverse shell code with provided `host` ip address and `port`. Useful for e.g. CTF and HackTheBox challenges.

## Usage

```bash
rsg --host 10.168.13.20 --port 4444
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
./rsg --host 127.0.0.1 --port 4444
```
