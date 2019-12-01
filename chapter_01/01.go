package main

import "fmt"

func main() {
  word := []rune("パタトクカシーー")
  for i := 1; i <= len(word)-1; i+=2 {
    fmt.Print(string(word[i]))
  }
}