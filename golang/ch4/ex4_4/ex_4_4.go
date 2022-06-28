package main

import "fmt"

func main() {
	a := 3	// int
	var b float64 = 3.5

	var c = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(a, b, c, d, e, f, g, h)
	// 3 3.5 3 9 7 63 10 9
}