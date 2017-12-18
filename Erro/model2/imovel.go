package model2

import "errors"

//Imovel é um struct que armazena dados de um imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

/*
ErrValorDoImovelInvalido é um erro apresentado quando é atribuido a imóvel com valor baixo.
*/
var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel")

// ErrValorDeImovelMuitoAlto  é um erro que é apresentado quando o valor do imóvel é muito alto.
var ErrValorDeImovelMuitoAlto = errors.New("O valor é muito alto.")

//GetValor atribui valor a variavel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}
