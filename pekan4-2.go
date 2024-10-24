package main

import "fmt"

func fibonacci(n int) int {
    if n == 0 || n == 1 {
        return n
    }

    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    var n, sn int
    fmt.Scanln(&n)
	sn = 0

    for i := 0; i <= n; i++ {
        sn = fibonacci(i)
    }
	fmt.Println(sn)
}
