package generator

import (
	"fmt"
	"math"
)

func FormatMoney(castka float64) string {
	castka = math.Floor(castka * 100)
	castka = math.Abs(castka)
	return fmt.Sprintf("%.0f", castka)
}
