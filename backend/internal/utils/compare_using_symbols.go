package utils

import (
	"fmt"
	"log"
)

func CompareUsingSymbol(symbol string, a float64, b float64) (bool, error) {
    log.Printf("Comparing: %f %s %f", a, symbol, b)
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
