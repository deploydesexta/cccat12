package cpf

import (
	"fmt"
	"strings"
)

type Cpf struct {
	value string
}

func NewCpf(value string) (Cpf, error) {
	if !validate(value) {
		return Cpf{}, fmt.Errorf("invalid cpf")
	}
	return Cpf{value: value}, nil
}

func validate(str string) bool {
	cpf := clean(str)
	if !hasValidLength(cpf) {
		return false
	}

	if hasSameDigits(cpf) {
		return false
	}

	dg1 := calculateDigit(cpf, 10)
	dg2 := calculateDigit(cpf, 11)
	return extractCheckDigit(cpf) == fmt.Sprintf("%d%d", dg1, dg2)
}

func extractCheckDigit(cpf string) string {
	return cpf[len(cpf)-2:]
}

func calculateDigit(cpf string, factor int) int {
	total := 0
	for _, digit := range cpf {
		if factor > 1 {
			total += int(digit-'0') * factor
			factor--
		}
	}
	rest := total % 11
	if rest < 2 {
		return 0
	} else {
		return 11 - rest
	}
}

func hasSameDigits(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

func hasValidLength(cpf string) bool {
	return len(cpf) >= 11 && len(cpf) <= 14
}

func clean(str string) string {
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	return str
}
