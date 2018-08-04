# gowk
CLI tool of golang similar to awk

## Installation

```
go get github.com/mitubaEX/gowk
```

## Usage

### help

```
❯❯❯ ./gowk -h

Usage: ./gowk <command> [<options>] <file>...
Options
  -c string
        target column number
  -d string
        delimiter for line (default ",")
  -f string
        function of filter condition (default "> 0")
```

### Sum

```
❯❯❯ echo "1,2,3\n2,3,4" | go run main.go sum
3
❯❯❯ echo "1,2,3\n2,3,4" | go run main.go sum -c 1,2
5,7
❯❯❯ echo "1,2,hello,3\n2,3,world,3" | go run main.go sum -c 2,3
helloworld,6
```

### filter

```
❯❯❯ echo "1,10,hello,3\n2,20,world,4"
1,10,hello,3
2,20,world,4

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | go run main.go filter -c 1,3 -f '>= 10.0'
10,0
20,0

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | go run main.go filter -c 1,3 -f '>= 10.1'
20,0

# string
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | go run main.go filter -c 2 -f '> "w"'
world
```

### frequency

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | go run main.go frequency -c 2
hello: 1
world: 1
```

### min, max

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | go run main.go max -c 3
4
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | go run main.go min -c 3
3
```
