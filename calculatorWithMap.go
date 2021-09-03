package calculatorWithMap

type Operation interface {
	Calculate(operand1, operand2 int) int
}

type AdditionService struct{}
type SubstractService struct{}
type MultipleService struct{}
type DivisionService struct{}

func (a AdditionService) Calculate(operand1, operand2 int) int {
	return operand1 + operand2
}

func (s SubstractService) Calculate(operand1, operand2 int) int {
	return operand1 - operand2
}

func (m MultipleService) Calculate(operand1, operand2 int) int {
	return operand1 * operand2
}

func (d DivisionService) Calculate(operand1, operand2 int) int {
	return operand1 / operand2
}

func NewCalculator() map[string]Operation {
	services := make(map[string]Operation)
	services["add"] = AdditionService{}
	services["substract"] = SubstractService{}
	services["multiple"] = MultipleService{}
	services["division"] = DivisionService{}

	return services
}