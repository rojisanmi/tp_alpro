package main

import "fmt"

func pangkat(x, y int) int {
    if y == 0 {
        return 1
    }

    return x * pangkat(x, y-1)
}

func main() {
    var x, y, hasil int
    fmt.Scanln(&x, &y)

    hasil = pangkat(x, y)
    fmt.Println(hasil)
}