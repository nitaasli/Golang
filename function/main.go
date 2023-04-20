package main

import "fmt"

func nita(name string) (text string) {
	text = "Hallo " + name
	return
}

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

func add(x, y int) int {
	return x + y
}

func main() {
	// nita(("manusia"))
	// fmt.Println(nita("manusia"))
	// fmt.Println(add(2, 5))
	fmt.Println(addSub(2, 5))
}
