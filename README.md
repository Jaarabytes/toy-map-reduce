# wc-mapreduce-go

toy mapreduce wordcount implementation in go

It works similar to the Counter python library

Takes input of random words, strings, numbers and outlines output (an array of structured structs each containing the word and the number of appearances alongside it).

Input : 
```
skillissue
skillissue
a
ww
skillissue
a
ww
ww
a
b
a
a
```


Output :
```
[{a 5} {ww 3} {skillissue 3} {b 1}]
```

## Usage

```
cat 1test.txt 2test.txt | go run main.go
```

You are free to modify the test txt files btw.

## Installation

Prerequisites:
    - Golang installed

Steps:
- Clone the repository
- Enjoy yourself!

## Contributions

I don't want them. Thank you!
