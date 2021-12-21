package usecase

import "errors"

var (
	errorAccountCPFExists = errors.New("já existe uma conta cadastrada com este CPF")
)
