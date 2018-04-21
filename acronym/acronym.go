package acronym

import (
	"strings"
)

func Abbreviate(text string) string {
	var sigla string

	texts := strings.Split(text, " ")
	for _, value := range texts {
		if strings.Contains(value, "-") {
			sigla += processar(strings.Split(value, "-"))
		} else {
			sigla += retornaPrimeiraLetra(value)
		}
	}
	return sigla
}

func processar(values []string) string {
	var sigla string
	for _, split := range values {
		sigla += retornaPrimeiraLetra(split)
	}
	return sigla
}

func retornaPrimeiraLetra(text string) string {
	return strings.ToUpper(string(text[0]))
}
