package solver

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type MathSolverStub struct {
}

func (ms MathSolverStub) Resolve(ctx context.Context, expression string) (float64, error) {
	switch expression {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}

func TestProcessorExpression(t *testing.T) {
	p := Processor{Solver: MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
(2 + 2 * 10`)
	data := []float64{22, 40, 0}
	for _, d := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		if err != nil {
			t.Error(err)
		}
		if result != d {
			t.Errorf("exprected result %f, got %f", d, result)
		}
	}
}
