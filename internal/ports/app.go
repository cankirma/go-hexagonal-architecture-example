package ports

type APIPort interface {
	GetAddition(a, b int32) (int32, error)
	GetSubtraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivision(a, b int32) (int32, error)
}

type Adapter struct {
	arith ArithmeticPort
	db    DbPort
}

func NewAdapter(arith ArithmeticPort, db DbPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}
func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)

	if err != nil {
		return 0, err

	}
	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)

	if err != nil {
		return 0, err

	}
	err = apia.db.AddToHistory(answer, "substraction")
	if err != nil {
		return 0, err

	}
	return answer, nil
}
func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)

	if err != nil {
		return 0, err

	}
	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err

	}
	return answer, nil
}
func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)

	if err != nil {
		return 0, err

	}
	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err

	}
	return answer, nil
}
