package main
import "fmt"
var a int

func procB(b, c *int) {
	*b = *b + *c
	*c = a + *b + *c
}

func main() {
	a = 10
	procB(&a, &a)
	fmt.Print(a)
}