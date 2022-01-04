package utils

import "regexp"

func TrimCPF(cpf string) (result string) {
	regex := regexp.MustCompile("[^0-9]+")
	result = regex.ReplaceAllString(cpf, "")
	return result
}
