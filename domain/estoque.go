package domain

type Estoque struct {
	ID         int64
	Quantidade int64
	Version    int64
}

func (e *Estoque) PodeComprar() bool {
	return e.Quantidade > 0
}
