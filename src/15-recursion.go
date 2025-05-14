package main

import "fmt"

func main(){
    countToZero(5)

    fmt.Println(fact(7))

    var fib func(n int) int

    fib = func(n int) int {
        if  n < 2 {
            return n
        }
        return fib(n-1)+fib(n-2)
    }

    fmt.Println(fib(7))
}

func countToZero(start int){
    if start <= 0 {
        return
    }
    fmt.Println(start)
    countToZero(start-1)
    return 
} 

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
