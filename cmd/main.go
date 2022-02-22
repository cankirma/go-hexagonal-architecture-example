package main

import (
	"fmt"
	"github.com/cankirma/go-hexagonal-architecture-example/internal/adapters/core/arithmetic"
)

func main() {

	arithAdapter := arithmetic.NewAdapter()
	res, _ := arithAdapter.Division(1123, 3)
	fmt.Println(res)
}
