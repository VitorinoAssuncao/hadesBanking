package validations

import (
	"testing"
)

func Test_nameIsNotEmpty(t *testing.T) {
	t.Run("Nome não etá vazio", func(t *testing.T) {
		name := "Joao"
		result := nameIsNotEmpty(name)

		if !result {
			t.Errorf("Nome não pode ser vazio")
		}
	})

	t.Run("Nome está vazio", func(t *testing.T) {
		name := ""
		result := nameIsNotEmpty(name)

		if result {
			t.Errorf("Nome não pode ser vazio")
		}
	})
}

func Test_cpfIsNotEmpty(t *testing.T) {
	t.Run("CPF não é vazia", func(t *testing.T) {
		cpf := "123456789"
		result := cpfIsNotEmpty(cpf)

		if !result {
			t.Errorf("CPF não pode ser vazio")
		}
	})

	t.Run("CPF está vazia", func(t *testing.T) {
		cpf := ""
		result := cpfIsNotEmpty(cpf)

		if result {
			t.Errorf("CPF não poderá ser vazia")
		}
	})
}

func Test_cpfIsJustNumber(t *testing.T) {
	t.Run("CPF é apenas números", func(t *testing.T) {
		cpf := "12345678912"
		result := cpfIsJustNumbers(cpf)

		if !result {
			t.Errorf("CPF deverá conter apenas números")
		}
	})

	t.Run("CPF contem simbolos", func(t *testing.T) {
		cpf := "123.456.789-12"
		result := cpfIsJustNumbers(cpf)

		if result {
			t.Errorf("CPF deverá conter apenas números")
		}
	})
}

func Test_secretIsNotEmpty(t *testing.T) {
	t.Run("Senha não é vazia", func(t *testing.T) {
		secret := "12345"
		result := secretIsNotEmpty(secret)

		if !result {
			t.Errorf("Senha não pode ser vazia")
		}
	})

	t.Run("Senha está vazia", func(t *testing.T) {
		secret := ""
		result := secretIsNotEmpty(secret)

		if result {
			t.Errorf("Senha não poderá ser vazia")
		}
	})
}

func Test_balanceIsPositive(t *testing.T) {
	t.Run("Realizando o teste com um valor de Saldo positivo", func(t *testing.T) {
		balance := 10
		result := balanceIsPositive(balance)

		if !result {
			t.Errorf("Saldo deve ser positivo ou igual a zero")
		}
	})

	t.Run("Realizando o teste com um valor de Saldo negativo", func(t *testing.T) {
		balance := -10
		result := balanceIsPositive(balance)

		if result {
			t.Errorf("Saldo só pode ser maior ou igual a zero")
		}
	})

	t.Run("Realizando o teste com o valor igual a zero", func(t *testing.T) {
		balance := 0
		result := balanceIsPositive(balance)

		if !result {
			t.Errorf("Saldo pode ser igual a zero")
		}
	})
}

func Test_calculateFirstVerifyingDigit(t *testing.T) {
	cpf := "383433358"
	cpfArray := make([]int, 0)

	for _, i := range cpf {
		value := int(i) - '0' //subtração devolve o valor do inteiro corretamente (49 - valor de 1, - valor do 0 (48))
		cpfArray = append(cpfArray, value)
	}

	firstDigit := calculateFirstVerifyingDigit(cpfArray)
	if firstDigit != 1 {
		t.Errorf("Calculo do digito verificador incorreto")
	}
}

func Test_calculateSecondVerifyingDigit(t *testing.T) {
	cpf := "3834333581"
	cpfArray := make([]int, 0)

	for _, i := range cpf {
		value := int(i) - '0' //subtração devolve o valor do inteiro corretamente (49 - valor de 1, - valor do 0 (48))
		cpfArray = append(cpfArray, value)
	}

	firstDigit := calculateSecondVerifyingDigit(cpfArray)
	if firstDigit != 3 {
		t.Errorf("Calculo do digito verificador incorreto")
	}
}

func Test_cpfIsValid(t *testing.T) {
	cpf := "38343335813"
	result := cpfIsValid(cpf)
	if !result {
		t.Errorf("Valor calculado inválido")
	}
}
