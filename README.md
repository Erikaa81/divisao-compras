# Divisão de Compras 🛒

## Objetivo
Receber uma lista de compras com: Itens, Quantidade de cada item, Preço por unidade/peso/pacote de cada item.
Em seguida dividir entre emails (cada email, representa uma pessoa)


### Regra Geral
- Como a menor unidade na nossa moeda é o centavo, use valores inteiros em centavos. Assim, um real será representado por 100 (cem centavos é igual a um real).


#### Função
- Calcula a soma dos valores, ou seja, multiplica o preço de cada item por sua quantidade e soma todos os itens;
- Divide o valor de forma igual entre a quantidade de e-mails;
- Caso a divisão nao for exata o resto é distribuido entre os e-mails, até nao sobrar nada;
- Retorna um mapa/dicionário onde a chave é o e-mail e o valor é quanto ele deve pagar nessa conta.


##### Sucesso 
    itens := []cesta.ListaMercado{{"cafe", 2, 2000}, {"pão", 2, 500}}
	emails := []string{"erika@hotmail.com", "joao@hotmail.com"}
    map[erika@hotmail.com:2500 joao@hotmail.com:2500]


###### Erros
- Não aceita lista vazia de emails;
- A inculsão de pelo menos um item na cesta é obrigatório.

(Teste!!)

