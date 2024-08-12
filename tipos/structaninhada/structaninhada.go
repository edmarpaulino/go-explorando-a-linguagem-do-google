package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() (total float64) {
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
}
