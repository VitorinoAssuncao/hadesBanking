package types

import "regexp"

type Document string

func (cpf Document) TrimCPF() (result Document) {
	regex := regexp.MustCompile("[^0-9]+")
	value := regex.ReplaceAllString(string(cpf), "")
	result = Document(value)
	return result
}

func (cpf Document) ToString() string {
	return string(cpf)
}
