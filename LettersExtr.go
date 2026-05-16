package main 
import "fmt"

func main() {
  name := "Hello World"
  
  letters := make(map[rune]int)
  
  for _, letter := range name {
     letters[letter]++
     fmt.Println(string(letter))
  }
  
}
