package model

//Imovel é um struct que armazena dados de um imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor atribui valor a variavel
func (i *Imovel) GetValor() int {
	return i.valor
}
