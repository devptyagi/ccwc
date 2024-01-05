## 'wc' Unix Tool written in Go

This tool is build as a solution to John Crickett's Coding Challenges.    
(https://codingchallenges.fyi/challenges/challenge-wc)

### Usage:

```
Usage: ccwc [options] <file>

Options:
  -c    Count bytes in the file
  -l    Count lines in the file
  -m    Count characters in the file
  -w    Count words in the file
  --help    Show this help message

If no file is provided, ccwc reads from stdin.

If no option is specified, ccwc defaults to -l -w and -c flags all together.
```

### Build:

```
go build -o ccwc ./cmd/ccwc  
```

### Test:

```
> ./ccwc -c test.txt
342190 test.txt
```

```
> ./ccwc -l test.txt
7145 test.txt
```

```
> ./ccwc -w test.txt
58164 test.txt
```

```
> ./ccwc -m test.txt
339292 test.txt
```

**Note**: -m output is considering UTF-8 encoding.

```
> ./ccwc test.txt
7145 58164 342190 test.txt
```

```
> cat test.txt | ./ccwc -l
7145
```