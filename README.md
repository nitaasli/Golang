# Golang

//code basic menampilkan atau print kata hello world

package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

//Level2

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

TURUN TIPE
complex : digunakan untuk tipe data yang memerlukan angka nilai imaginer
error : adalah integface tipe dari string yang digunakan untuk menangani error di golang 
interface : digunakan untuk penentu metode dan menyimpan data kosong

ATURAN FUNCTION AT GO-LANG :
1. Tidak bisa dimulai dengan ANGKA
2. Tidak dapat memiliki SPASI
3. Tidak bisa dimulai dengan kata SPECIAL kecuali dengan tanda (_)

