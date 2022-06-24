package main

import "log"

func main() {
	log.Println("hello world")

	for i := range 10 {
		log.Println(i)
	}

	log.Println("Finished")

	log.Println("sum 1+2 = ", sum(1, 2))
	log.Println("sub 1-2 = ", sub(1, 2))
	log.Println("mul 1-2 = ", mul(1, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a * b
}

// tdsfs
// sdfs
