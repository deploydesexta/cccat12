package cpf

import (
	"fmt"
	"strings"
)

func Validate(str string) bool {
	if str != "" {
		if len(str) >= 11 && len(str) <= 14 {
			str = strings.ReplaceAll(str, ".", "")
			str = strings.ReplaceAll(str, "-", "")
			str = strings.ReplaceAll(str, " ", "")

			if !strings.HasPrefix(strings.Repeat(string(str[0]), len(str)), str) {
				var d1, d2, dg1, dg2, rest, digito int
				d1, d2, dg1, dg2, rest = 0, 0, 0, 0, 0

				for nCount := 1; nCount < len(str)-1; nCount++ {
					digito = int(str[nCount]) - '0'
					d1 = d1 + (11-nCount)*digito
					d2 = d2 + (12-nCount)*digito
				}

				rest = d1 % 11
				dg1 = 0
				if rest >= 2 {
					dg1 = 11 - rest
				}
				d2 += 2 * dg1
				rest = (d2 % 11)
				dg2 = 0
				if rest >= 2 {
					dg2 = 11 - rest
				}

				nDigVerific := str[len(str)-2:]
				nDigResult := fmt.Sprintf("%d%d", dg1, dg2)

				return nDigVerific == nDigResult
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
