
package main

import "testing"

func TestSubtrair(t *testing.T) {
	resultado := subtrair(10, 5)
	esperado := 5
	
	if resultado != esperado {
		t.Errorf("Erro ao subtrair: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestSubtrairNegativo(t *testing.T) {
	resultado := subtrair(5, 10)
	esperado := -5
	
	if resultado != esperado {
		t.Errorf("Erro ao subtrair números negativos: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestMultiplicar(t *testing.T) {
	resultado := multiplicar(10, 5)
	esperado := 50
	
	if resultado != esperado {
		t.Errorf("Erro ao multiplicar: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestMultiplicarPorZero(t *testing.T) {
	resultado := multiplicar(10, 0)
	esperado := 0
	
	if resultado != esperado {
		t.Errorf("Erro ao multiplicar por zero: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestMultiplicarNegativo(t *testing.T) {
	resultado := multiplicar(-5, 3)
	esperado := -15
	
	if resultado != esperado {
		t.Errorf("Erro ao multiplicar com negativo: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestDividir(t *testing.T) {
	resultado := dividir(10, 5)
	esperado := 2
	
	if resultado != esperado {
		t.Errorf("Erro ao dividir: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestDividirPorUm(t *testing.T) {
	resultado := dividir(10, 1)
	esperado := 10
	
	if resultado != esperado {
		t.Errorf("Erro ao dividir por 1: esperado %d, obtido %d", esperado, resultado)
	}
}

func TestDividirNegativo(t *testing.T) {
	resultado := dividir(-10, 5)
	esperado := -2
	
	if resultado != esperado {
		t.Errorf("Erro ao dividir com negativo: esperado %d, obtido %d", esperado, resultado)
	}
}