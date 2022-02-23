package api

import "github.com/cankirma/go-hexagonal-architecture-example/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}
