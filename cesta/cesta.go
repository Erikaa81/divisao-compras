package cesta

import (
	"errors"
	"reflect"
)

type ListaMercado struct {
	Item           string
	QuantidadeItem int
	Valor          int
}

func DivisaoPorPessoa(emails []string, cesta []ListaMercado) (map[string]int, error) {
	var total int
	divisaoFinal := make(map[string]int)

	for item := range emails {
		if emails[item] == "" {
			return nil, errors.New("Email é obrigatório")
		}

		if reflect.DeepEqual(cesta, []ListaMercado{}) == true {
			return nil, errors.New("É necessario ter pelo menos um item na cesta")

		}

	}

	for _, item := range cesta {
		custoTotal := item.QuantidadeItem * item.Valor
		total += custoTotal

	}

	custoTotal := total / len(emails)
	resto := total % len(emails)

	for _, DivisaoPorPessoa := range emails {
		divisaoFinal[DivisaoPorPessoa] += custoTotal

		if custoTotal*len(emails) < total {
			divisaoFinal[DivisaoPorPessoa] += resto
			resto -= resto
		}

	}

	return divisaoFinal, nil
}
