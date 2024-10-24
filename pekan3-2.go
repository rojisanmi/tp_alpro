package main
import "fmt"

func hitungReamur(c float64, r *float64) {
	*r = c * 4 / 5
}

func hitungFahrenheit(c float64, f *float64) {
	*f = (c * 9 / 5) + 32
}

func hitungKelvin(c float64, k *float64) {
	*k = c + 273.15
}

func main() {
	var celsius, reamur, fahrenheit, kelvin float64
	fmt.Scan(&celsius)
	
	hitungReamur(celsius, &reamur)
	hitungFahrenheit(celsius, &fahrenheit)
	hitungKelvin(celsius, &kelvin)
	fmt.Printf("%.fR %.fF %.2fK", reamur, fahrenheit, kelvin)
}