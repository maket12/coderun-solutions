package main

import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1<<20)

  line, _ := reader.ReadString('\n')
  line = strings.TrimSpace(line)

  words := strings.Fields(line)

  var numbers []int = make([]int,0, 3)
  for _, word := range words {
    num, _ := strconv.Atoi(word)
    numbers = append(numbers, num)
  }

  for i := 0; i < len(numbers) - 1; i++ {
    for j := i + 1; j < len(numbers); j++ {
      if numbers[i] > numbers[j] {
        numbers[i], numbers[j] = numbers[j], numbers[i]
      }
    }
  }

  writer := bufio.NewWriterSize(os.Stdout, 1<<20)
  defer writer.Flush()

  writer.WriteString(strconv.Itoa(numbers[1]))
  writer.WriteByte('\n')
}