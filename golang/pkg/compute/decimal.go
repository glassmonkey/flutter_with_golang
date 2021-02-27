package compute

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func initValues(v1, v2 string) (decimal.Decimal, decimal.Decimal, error) {
	uv1, err:= decimal.NewFromString(v1)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, fmt.Errorf("failed initilize first value, because %s", err)
	}

	uv2, err:= decimal.NewFromString(v2)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, fmt.Errorf("failed initilize second value, because %s", err)
	}
	return uv1, uv2, nil
}

func Add(v1, v2 string) (string, error) {
	uv1, uv2, err := initValues(v1, v2)
	if err != nil {
		return "", err
	}

	return uv1.Add(uv2).String(), nil
}

func Sub(v1, v2 string) (string, error) {
	uv1, uv2, err := initValues(v1, v2)
	if err != nil {
		return "", err
	}

	return uv1.Sub(uv2).String(), nil
}


