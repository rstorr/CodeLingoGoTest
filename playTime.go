package main

import "fmt"

type myVertex struct {
	X int
	Y int
}

func structFunc() {
	v := myVertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func add(a int, b int) int {
	return a + b
}

func addLoop(a int, b int) int {
	sum := a
	for i := 0; i < b; i++ {
		sum += i
	}
	return sum
}

func arrayFunc() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "ontoborn"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slicesFunc() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	s := primes[1:4]

	fmt.Println(s)
}

func main() {
	fmt.Println("sum is ", add(10, 19))
	fmt.Println("Total sum =", addLoop(0, 8))
	arrayFunc()
	slicesFunc()
	structFunc()
}
