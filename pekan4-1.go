package main

import "fmt"

func cetakGanjil(n int) {
	if n % 2 == 0 {
		n -= 1
	}
	
	if n == -1 {
        return
    }

    cetakGanjil(n - 1)
    fmt.Print(n, " ")
}

func main() {
    var n int
    fmt.Scanln(&n)
	
    cetakGanjil(n)
}