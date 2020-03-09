package model

import (
	"errors"
	"fmt"
)

/*
Imovel -> representa informações de um imóvel
*/
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int    `json: valor`
}

// ErrValorDeImoveInvalido é um erro que é aprensentado quanto é atribuido a imóvel um valor muito baixo
var ErrValorDeImoveInvalido = errors.New("O valor informado não é válido para um imovél")

/*
SetValor -> define o valor do imóvel
*/
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImoveInvalido
		fmt.Println(err)
		return
	}
	i.valor = valor
	return
}

/*
GetValor -> retorna o valor do imóvel
*/
func (i *Imovel) GetValor() int {
	return i.valor
}
