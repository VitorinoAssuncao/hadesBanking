package transfer

import (
	"stoneBanking/app/domain/types"
	"time"
)

type Transfer struct {
	//ID único da transfêrencia (Númerico gerado pelo banco)
	ID int

	External_ID types.TransferID

	//Código da conta de origem
	Account_origin_id types.AccountOriginID

	//Código da conta de destino
	Account_destination_id types.AccountDestinyID

	//Valor referente a transfêrencia, em centavos brasileiros (R$)
	Amount types.Money

	//Data da transfêrencia
	Created_at time.Time
}
