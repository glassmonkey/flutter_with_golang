package api

import (
	"fmt"
	"reflect"

	"github.com/Knetic/govaluate"
)

func GoCall() (string, error) {
	return "hello, world by golang", nil
}

func Evaluate(formula string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return 0, err
	}
	i := make(map[string]interface{})

	result, err := expression.Evaluate(i)
	if err != nil {
		return 0, err
	}
	v, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid formula result. formula[%s], result[%#v](%s)", formula, result, reflect.TypeOf(result))
	}
	return v, nil
}
