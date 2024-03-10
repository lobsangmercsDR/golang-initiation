// Package declaration
package main

// Importing required packages
import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
)

var baseCuadrado int
var radius float64
var baseTriangulo int
var alturaTriangulo int
var baseTrapecio int


func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: 
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func areaCuadrado(base int) {

	area := base * base
	fmt.Printf("El area del cuadrado es:  %.v \n", area)

}

func areaCirculo(radius float64) {
	PI := math.Pi
	AreaCirculo := PI * radius * radius
	fmt.Printf("El area del circulo es: %.3v  \n	", AreaCirculo)
}

func areaTriangulo(base, altura int) {
	area := base * altura / 2
	fmt.Printf("El area del triangulo es: %.v  \n", area)
}
func main() {

	var name string

	println("Calculadora Geometrica")
	println("Ingresa tu nombre")
	fmt.Scanf("%s", &name)
	fmt.Printf("Bienvenido %s gracias por usar mi calcuculadora geometrica \n", name)

	println("Selecciona la operacion que deseas realizar")
	println("1.Calcular area de cuadrado")
	println("2.Calcular area de ciruclo")
	println("3.Calcular area de triangulo")

	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		println("Ingrese la base del cuadrado")
		fmt.Scan(&baseCuadrado)
		areaCuadrado(baseCuadrado)
		println("Desea realizar otra Operacion?")
		println("1.Si")
		println("2.No")
		var Resp int
		fmt.Scan(&Resp)
		switch Resp {
		case 1:
			clearConsole()
			main()
		case 2:
			os.Exit(0)
		}

	case 2:
		println("Ingrese el radio del circulo")
		fmt.Scan(&radius)
		areaCirculo(radius)
		println("Desea realizar otra Operacion?")
		println("1.Si")
		println("2.No")
		var Resp int
		fmt.Scan(&Resp)
		switch Resp {
		case 1:
			clearConsole()
			main()
		case 2:
			os.Exit(0)
		}
	case 3:
		println("Ingrese la base del triangulo")
		fmt.Scan(&baseTriangulo)
		println("Ingrese la altura del triangulo")
		fmt.Scan(&alturaTriangulo)
		areaTriangulo(baseTriangulo, alturaTriangulo)
		println("Desea realizar otra Operacion?")
		println("1.Si")
		println("2.No")
		var Resp int
		fmt.Scan(&Resp)
		switch Resp {
		case 1:
			clearConsole()
			main()
		case 2:
			os.Exit(0)
		}

	default:
		println("Opcion no valida")
		fmt.Println("Presiona enter para continuar...")
		fmt.Scanln()

		main()

	}

}
