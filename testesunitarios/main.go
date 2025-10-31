package main

func subtrair(a int, b int) int {
	return a - b
}

func multiplicar(a int, b int) int {
	return a * b
}

func dividir(a int, b int) int {
	return a / b
}

func main() {
	println("Subtração: ", subtrair(10, 5))
	println("Multiplicação: ", multiplicar(10, 5))
	println("Divisão: ", dividir(10, 5))
}

