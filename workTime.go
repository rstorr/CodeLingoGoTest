/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

func reubenNewFunc() {
	s := "stringA"
	s1 := "stringB"
	fmt.Println(s + s1)
}

func addSomeNums() {
	numA := 7
	numB := 5
	fmt.Println(numA + numB)
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

func imCoolFunc(str string) {
	fmt.Println(str + " is cool.")
}

func main() {
	fmt.Println("sum is ", add(10, 19))
	fmt.Println("Total sum =", addLoop(0, 8))
	arrayFunc()
	slicesFunc()
	structFunc()
	reubenNewFunc()
	addSomeNums()
	imCoolFunc("Reuben")
}
