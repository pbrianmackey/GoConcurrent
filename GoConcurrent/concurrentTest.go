package main

import (
  "fmt"
  "time"
)

func main() {
  go say("Thread 1 complete!", 1*time.Second)
  go say("Thread 2 complete!", 2*time.Second)
  go say("Thread 3 complete!", 3*time.Second)
  time.Sleep(4 * time.Second)
}

func say(text string, delay time.Duration){
    time.Sleep(delay)
    fmt.Println(text)
}
