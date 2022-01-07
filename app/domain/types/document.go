package types

import "regexp"

type Document string

func (cpf Document) TrimCPF() (result string) {
	regex := regexp.MustCompile("[^0-9]+")
	result = regex.ReplaceAllString(string(cpf), "")
	return result
}

func (cpf Document) ToString() string {
	return string(cpf)
}
