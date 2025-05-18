package main

import "fmt"

func main() {
    nums := []int{2,4,6}

    sum := 0

    for _, num := range nums {
        sum += num
    }

    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 4  {
            fmt.Println("4 is at idx:", i)
        }
    }

    kvs := map[string]string{"a":"apple", "b":"bananananna"}
    for k, v := range kvs {
        fmt.Printf("%s --> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key", k)
    }

    for i, c := range "golang" {
        fmt.Println(i,c)
    }
}
