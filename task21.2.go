package main

import "fmt"

func office(A func(int, int) int) (result int) {
	defer func() {
		result = A(10, 5)
	}()
	return 0
}

func main() {

	fmt.Println("21.6 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 2. Анонимные функции")
	fmt.Println("------------")

	a := func(x, y int) int { return x + y }
	b := func(x, y int) int { return x * y }
	c := func(x, y int) int { return x / y }

	fmt.Println(office(a))
	fmt.Println(office(b))
	fmt.Println(office(c))
}
