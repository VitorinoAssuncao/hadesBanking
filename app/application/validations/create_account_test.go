package validations

import "testing"

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
