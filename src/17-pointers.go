package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1

    fmt.Println("initial", i)

    // This should not change anything, as the parameter value is passed not the reference
    zeroval(i)
    fmt.Println("zeroval", i)

    // Since we are passing the address and updating the value this should print 0
    zeroptr(&i)
    fmt.Println("zeroptr", i)

    fmt.Println("pointer", &i)
}
