# gowk
CLI tool of golang similar to awk

## Installation

```
go get github.com/mitubaEX/gowk
```

## Usage

### help

```
❯❯❯ gowk -h

Usage: gowk <command> [<options>] <file>...
Commands:
        sum
        filter
        frequency
        min
        max
        length
        intersection

Options:
  -c string
        target column number (default "0")
  -d string
        delimiter for line (default ",")
  -f string
        function of filter condition {required at filter mode} (default "> 0")
```

### Sum

Get sum in per column.

```
❯❯❯ echo "1,2,3\n2,3,4" | gowk sum
3
❯❯❯ echo "1,2,3\n2,3,4" | gowk sum -c 1,2
5,7
❯❯❯ echo "1,2,hello,3\n2,3,world,3" | gowk sum -c 2,3
helloworld,6
```

### filter

Filter given column by given condition.

```
❯❯❯ echo "1,10,hello,3\n2,20,world,4"
1,10,hello,3
2,20,world,4

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | gowk filter -c 1,3 -f '>= 10.0'
10,0
20,0

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | gowk filter -c 1,3 -f '>= 10.1'
20,0

# string
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk filter -c 2 -f '> "w"'
world
```

### frequency

Analyze given column word frequency.

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk frequency -c 2
hello: 1
world: 1
```

### min, max

Get per given column min, max.

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk max -c 3
4
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk min -c 3
3
```

### length

Count per line elements by given delimiter.

```
❯❯❯ echo "1,2,3,4\n1,2,3" | go run main.go length
4
3
```

### intersection

```
❯❯❯ echo "hello,world\nworld,hello\nhello,world" | go run main.go intersection -c 0,1
hello
world
```
