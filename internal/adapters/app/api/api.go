package api

import "github.com/cankirma/go-hexagonal-architecture-example/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DbPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{
		arith: arith,
		db:    db,
	}
}
