package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true 
}

func main() {
    done := make(chan bool, 1)
    go worker(done)

    // Wait until we receive the message that the function is done running
    <-done
}
