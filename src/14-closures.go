package main

import "fmt"

func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}

// Returns a function that returns an int
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
