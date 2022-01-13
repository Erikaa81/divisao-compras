package cesta

import (
	"reflect"
	"testing"
)

type operandos struct {
	itens  []ListaMercado
	emails []string
}

func TestDivisaoPorPessoa(t *testing.T) {

	testes := []struct {
		nome         string
		operandos    operandos
		esperado     map[string]int
		erroEsperado bool
	}{
		{
			nome: "Divisao exata",
			operandos: operandos{

				[]ListaMercado{{"cafe", 2, 2000}, {"pão", 2, 500}},
				[]string{"erika@hotmail.com", "joao@hotmail.com"},
			},
			esperado: map[string]int{"erika@hotmail.com": 2500, "joao@hotmail.com": 2500},
		},
		{
			nome: "Divisao com resto",
			operandos: operandos{
				[]ListaMercado{{"cafe", 2, 2000}, {"pão", 2, 500}},
				[]string{"erika@hotmail.com", "joao@hotmail.com", "paulo@hotmail.com"},
			},
			esperado: map[string]int{"erika@hotmail.com": 1667, "joao@hotmail.com": 1667, "paulo@hotmail.com": 1666},
		},

		{
			nome: "Obrigatoriedade de pelo menos um item na cesta",
			operandos: operandos{

				[]ListaMercado{},
				[]string{"erika@hotmail.com", "joao@hotmail.com"},
			},
			erroEsperado: true,
		},

		{
			nome: "Email obrigatório ",
			operandos: operandos{

				[]ListaMercado{{"cafe", 2, 2000}},
				[]string{""},
			},
			erroEsperado: true,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {

			obtive, err := DivisaoPorPessoa(tt.operandos.emails, tt.operandos.itens)

			if (err != nil) != tt.erroEsperado {
				t.Errorf("DivisaoPorPessoa() erro : %v, erroEsperado: %v", err, tt.erroEsperado)
				return
			}

			if !reflect.DeepEqual(obtive, tt.esperado) {
				t.Errorf("DivisaoPorPessoa() : %v,  esperado %v", obtive, tt.esperado)
			}
		})
	}
}
