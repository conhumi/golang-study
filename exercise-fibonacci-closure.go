package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    num0 := 0
    num1 := 1
    return func() int {
        tmp := num0
        num0 = num1
        num1 = tmp + num1
        return tmp
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

