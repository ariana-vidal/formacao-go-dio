package main

func divisivelpor3() map[int]string {
	mapeamento := make(map[int]string)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			mapeamento[i] = "é divisível por 3"
		}
	}
	return mapeamento
}

func pinPan() map[int]string {
	mapeamento := make(map[int]string)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			mapeamento[i] = "Pin"
		} else if i%5 == 0 {
			mapeamento[i] = "Pon"
		}
	}
	return mapeamento
}

func main() {
	divisivelpor3 := divisivelpor3()
	for numero, descricao := range divisivelpor3 {
		println(numero, descricao)
	}

	pinPan := pinPan()
	for numero, descricao := range pinPan {
		println(numero, descricao)
	}

}