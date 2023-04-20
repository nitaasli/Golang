package main

import "fmt"

func main() {
	var variableName1 string = "Hello guys"
	variableName2 := "Hellow world"
	fmt.Println(variableName1)
	fmt.Println(variableName2)

	// primitive : boolean,int,float,string
	// boolean(nilai yang hanya menyimpan benar atau salah saja)

	boolVar := true
	fmt.Printf("Type: %T value: %v\n", boolVar, boolVar)

	//int

	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T value: %v\n", intVar2, intVar2)

}
