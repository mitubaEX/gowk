# gowk
CLI tool of golang for processing csv files.

## Installation

Use go get command.
Please run following command.

```
go get github.com/mitubaEX/gowk
```

## gowk subcommands

- [help](#help)
- [Sum](#sum)
- [filter](#filter)
- [frequency](#frequency)
- [min, max](#minmax)
- [length](#length)
- [intersection](#intersection)
- [distinct](#distinct)

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
        distinct

Options:
  -c string
        target column number (default "0")
  -d string
        delimiter for line (default ",")
  -f string
        function of filter condition {required at filter mode} (default "> 0")
  -v    set verbose mode (default false)
```

### Sum

Get sum in each column.

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

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | ./gowk filter -c 1 -f '>= 10.0'
1,10,hello,3
2,20,world,4

❯❯❯ echo "1,10,hello,3\n2,20,world,4" | gowk filter -c 1 -f '>= 10.1'
2,20,world,4

# string
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk filter -c 2 -f '> "w"'
2,3,world,4
```

### frequency

Analyze given column word frequency.

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk frequency -c 2
hello,1
world,1
```

### min,max

Get each given column min, max.

```
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk max -c 3
4
❯❯❯ echo "1,2,hello,3\n2,3,world,4" | gowk min -c 3
3
```

### length

Count each line elements by given delimiter.

```
❯❯❯ echo "1,2,3,4\n1,2,3" | gowk length
4
3
```

### intersection

Get intersection set.

```
❯❯❯ echo "hello,world\nworld,hello\nhello,world" | gowk intersection -c 0,1
hello
world
```

### distinct

Get distinct set.

```
❯❯❯ echo "1,world,hello,3\n2,3,world       ,4" | gowk distinct -c 1,2
3
hello
```

## verbose mode

In verbose mode, gowk print debug message.

```
❯❯❯ echo "1,2,3\n2,3,4" | gowk sum -v
2019/02/01 00:05:41 current element is 1
2019/02/01 00:05:41 current element is 2
3
```
