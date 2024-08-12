package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar mais métodos
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func soEsportivo(e esportivo) {
	e.ligarTurbo()
}

func soLuxuoso(l luxuoso) {
	l.fazerBaliza()
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()

	soEsportivo(b)
	soLuxuoso(b)

	fmt.Println(b)
}
