package main

// Formateo e impresión
import (
	"fmt"
	"math/rand/v2"
)

func suma(a int, b int) int {
	return a + b
}

func numeroRandom(min int, max int) int {
	return rand.IntN(max-min) + min
}

func main() {
	fmt.Println("¡Hola, mundo!") // Función principal
	var a int = numeroRandom(0, 10)
	var b int = numeroRandom(0, 10)
	fmt.Println("Número aleatorio:", a)
	fmt.Println("Número aleatorio:", b)
	fmt.Println("Suma:", suma(a, b))
}

//ejemplos de condicionales y blucles

// func main() {
//     // Condicional
//     edad := 20
//     if edad >= 18 {
//         fmt.Println("Eres mayor de edad.")
//     } else {
//         fmt.Println("Eres menor de edad.")
//     }

//     // Bucle for
//     for i := 0; i < 5; i++ {
//         fmt.Println("Iteración:", i)
//     }

//     // Switch
//     dia := "lunes"
//     switch dia {
//     case "lunes":
//         fmt.Println("Es lunes.")
//     case "viernes":
//         fmt.Println("Es viernes.")
//     default:
//         fmt.Println("Es otro día.")
//     }
// }
