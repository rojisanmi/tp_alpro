package main
import "fmt"

func mengurutkan(a, b *int) {
	var temp int
	if *b < *a {
		temp = *a
		*a = *b
		*b = temp
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	
	mengurutkan(&a, &b)
	fmt.Print(a, b)
}