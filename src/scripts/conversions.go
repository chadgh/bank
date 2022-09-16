package scripts

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertAmountToCents(amount string) (int32, error) {
	parts := strings.Split(amount, ".")
	dollars, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	cents, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	total := int32(cents + (dollars * 100))
	return total, nil
}

func ConvertCentsToAmount(cents int32) string {
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("%d.%.2d", dollars, remainder)
}
