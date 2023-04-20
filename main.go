package main

import (
	"errors"
	"fmt"
)

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

	//float
	floatVar1 := float32(3.2)
	floatVar2 := float64(3.2)
	fmt.Printf("Type: %T value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T value: %v\n", floatVar2, floatVar2)

	//bytes
	bytesVar := byte(255)
	fmt.Printf("Type: %T value: %v\n", bytesVar, bytesVar)

	bytesVar2 := []byte("Hello guys")
	fmt.Printf("Type: %T value: %v\n", bytesVar2, bytesVar2)

	//rune
	runeVar := '😄'
	fmt.Printf("Type: %T value: %v\n", runeVar, runeVar)

	//complex
	complexVar := -7 + 3i
	fmt.Printf("Type: %T value: %v\n", complexVar, complexVar)

	//error
	errVar := errors.New("error detected")
	fmt.Printf("Type: %T value: %v\n", errVar, errVar)

}
