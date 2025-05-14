package main

import "fmt"

func main() {
    sumOfNums := sum(1,2)
    fmt.Println(sumOfNums)
    sumOfNums = sum(1,2,3)
    fmt.Println(sumOfNums)

    nums := []int{1,2,3,4,5}
    sumOfNums = sum(nums...)
    fmt.Println(sumOfNums)
}

func sum(nums ...int) int {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }

    return total
}
