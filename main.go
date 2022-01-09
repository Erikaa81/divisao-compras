package main

import (
	"fmt"

	"github.com/Erikaa81/divisao-compras/cesta"
)

func main() {
	itens := []cesta.ListaMercado{{"cafe", 2, 2000}, {"pão", 2, 500}}
	emails := []string{"erika@hotmail.com", "joao@hotmail.com", "paulo@hotmail.com"}

	fmt.Println(cesta.DivisaoPorPessoa(emails, itens))

}
