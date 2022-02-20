package main

import "fmt"

func main() {

	fmt.Println("21.6 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Расчёт по формуле")
	fmt.Println("------------")

	var (
		x int16
		y uint8
		z float32
	)

	fmt.Println("Введите значение для x:")
	fmt.Scan(&x)
	fmt.Println("Введите значение для y:")
	fmt.Scan(&y)
	fmt.Println("Введите значение для z:")
	fmt.Scan(&z)

	S := func(x int16, y uint8, z float32) float32 { return 2*float32(x) + float32(y*y) - 3/z }
	fmt.Println(S(x, y, z))
}
