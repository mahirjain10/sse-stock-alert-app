package utils

import "fmt"

func CompareUsingSymbol(symbol string, a, b int64) (bool, error) {
	switch symbol {
	case "=":
		return a == b, nil
	case ">":
		return a > b, nil
	case "<":
		return a < b, nil
	case ">=":
		return a >= b, nil
	case "<=":
		return a <= b, nil
	default:
		return false, fmt.Errorf("unsupported symbol: %s", symbol)
	}
}