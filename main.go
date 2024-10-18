package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "sort"
	"strings"
	"sync"
)

func main () {
  scanner := bufio.NewScanner(os.Stdin)
  var input []string
  for scanner.Scan() {
    input = append(input, scanner.Text())
    if len(input) == 0 {
    }
  }

  mapped := make(chan map[string]int)
  result := make([]string, 0)
  var wg sync.WaitGroup

  for _, line := range input {
    wg.Add(1)
    go func (l string) {
      defer wg.Done()
      mapped <- mapFunc(l)
    }(line)
  }

  // Who doesn't love concurrency
  go func () {
     wg.Wait()
     close(mapped)
  }()

  reduced := reduceFunc(mapped)

  for word, count := range reduced {
    wordCountString := "{" + word + " " + strconv.Itoa(count) + "}"
    result = append(result, wordCountString)
  }
  sort.Strings(result)
  fmt.Println(result)
}

func mapFunc (input string) map[string]int {
  result := make(map[string]int)
  words := strings.Fields(input)
  for _, word := range words {
    result[strings.ToLower(word)]++
  }
  return result
}

func reduceFunc (mapped <- chan map[string]int) map[string]int {
    result := make(map[string]int)
    for m := range mapped {
    for word, count := range m {
        result[word] += count
    }
  }
  return result
}
