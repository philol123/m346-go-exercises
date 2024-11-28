package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]
	// TODO: correct up to index 4 using direct element access

	fibs = append(fibs, fibs[3]+fibs[4])
	fibs = append(fibs, fibs[4]+fibs[5])
	fibs = append(fibs, fibs[5]+fibs[6])
	fibs = append(fibs, fibs[6]+fibs[7])
	// TODO: compute three more Fibonacci numbers and append them

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
