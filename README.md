# gowk
CLI tool of golang similar to awk

## Installation

```
go get github.com/mitubaEX/gowk
```

## Usage

### Sum

```
❯❯❯ echo "1,2,3\n2,3,4" | go run main.go sum
3
❯❯❯ echo "1,2,3\n2,3,4" | go run main.go sum -c 1,2
5,7
```
