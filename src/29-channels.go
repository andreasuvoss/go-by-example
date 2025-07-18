package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages

    fmt.Println(msg)

    go func() { messages <- "pong" }()

    msg2 := <-messages

    fmt.Println(msg2)
}
