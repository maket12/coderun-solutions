package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)

  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)

  n, _ := strconv.Atoi(line)
}