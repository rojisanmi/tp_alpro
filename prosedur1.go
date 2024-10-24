package main

import "fmt"

const pi = 3.14

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(ll, lp, kl, kp float64, toLuas, totKel *float64) {
	*toLuas = ll + lp
	*totKel = kl + kp
}

func main() {
	var radius, sisi float64
	var ll, kl, lp, kp, toLuas, totKel float64
	var i int

	i = 1

	for {
		_, err := fmt.Scanln(&radius, &sisi)
		if err != nil {
			continue
		}

		if radius == 0 || sisi == 0 {
			break
		}

		if i == 1 {
			fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
		}

		i += 1

		hitungLuasKelilingLingkaran(radius, &ll, &kl)
		hitungLuasKelilingPersegi(sisi, &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &toLuas, &totKel)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", radius, sisi, ll, lp, kl, kp, toLuas, totKel)
	}
}
