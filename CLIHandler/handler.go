package CLIHandler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/calc"
)

type CLIHandler struct {
	calculator calc.Calculator
	stdout     io.Writer
}

func NewCLIHandler(calculator calc.Calculator, stdout io.Writer) *CLIHandler {
	return &CLIHandler{
		calculator: calculator,
		stdout:     stdout,
	}
}

func (this CLIHandler) Handler(args []string) error {
	if len(args) != 2 {
		return errors.New("Two arguments must be provided")
	}
	value1, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	value2, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	result := this.calculator.Calculate(value1, value2)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return err
	}
	return nil
}
