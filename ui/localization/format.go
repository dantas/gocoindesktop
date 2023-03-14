package localization

import "fmt"

func FormatPrice(price float64) string {
	return fmt.Sprintf("$ %.2f", price)
}
