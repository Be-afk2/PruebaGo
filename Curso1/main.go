package main

import (
	"fmt"
	"reflect"
)

func main() {
	//comentario uwu

	/*
		varias lineas de comentario
	*/
	// fmt.Print("Hola Mundo")

	var nombre string = "una cadena de texto"

	fmt.Println(string(nombre))

	var numero int = 10

	numero = numero + 1

	fmt.Println(numero)

	fmt.Println(nombre, numero)

	// imprimir el tipo de dato paquete reflect
	fmt.Println(reflect.TypeOf(nombre))

	var myfloat float32 = 3.14

	fmt.Println(myfloat)

	// al sumar 2 tipos de datos numerico distintos hay que expesificar el tipo de dato al momento de sumarlo
	// si es un float y un int trasformar uno de los 2 para que lo 2 sean el mismo tipo de dato
	fmt.Println(float32(numero) + myfloat)

	var mybool bool = true

	//variable declarada e inicializada en una sola linea
	myint := 10

	fmt.Println(myint)

	// const myconst int = 10

	// control de flujo

	if mybool || myint == 10 {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	// for _, i := range nombre {
	// 	fmt.Println(string(i))
	// }

	//array
	var myarray [3]int

	myarray[0] = 1
	myarray[1] = 2
	myarray[2] = 3
	fmt.Println(myarray)

	//Maps

	//Structs y anidacion de structs . Simulacion de class
	type Motor struct {
		Tipo     string
		Potencia int
	}

	type Auto struct {
		ID     int
		Modelo string
		Motor  Motor // `struct` anidado
	}

	auto := Auto{
		ID:     3,
		Modelo: "Model S",
		Motor: Motor{
			Tipo:     "Eléctrico",
			Potencia: 670,
		},
	}

	mymap := make(map[string]Auto)

	mymap["benja"] = auto
	mymap["pedro"] = auto
	mymap["juan"] = auto

	fmt.Println(mymap)
	fmt.Println(mymap["benja"])

	fmt.Println(auto)
}
