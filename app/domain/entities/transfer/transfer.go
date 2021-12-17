package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	//ID único da transfêrencia
	ID types.TransferID `json:"id"`

	//Código da conta de origem
	Account_origin_id types.AccountOriginID `json:"acount_origin_id"`

	//Código da conta de destino
	Account_destination_id types.AccountDestinyID `json:"acount_destination_id"`

	//Valor referente a transfêrencia, em centavos brasileiros (R$)
	Amount types.Money `json:"amount"`

	//Data da transfêrencia
	Created_at time.Time `json:"created_at"`
}
