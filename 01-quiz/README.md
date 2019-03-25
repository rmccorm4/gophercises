# 01-Quiz

This program reads questions/answers from a CSV file and makes
a timed quiz out of them.

## Usage

```bash
> go build main.go

> ./main -h

Usage of ./main:
  -filename string
    CSV file containing quiz questions/answers. (default "quiz.csv")
  -time int
    Time limit to finish the quiz in seconds. (default 30)
```


