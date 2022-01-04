package input

type CreateTransferVO struct {
	AccountOriginID  string `json:"account_origin_id" example:"2"`
	AccountDestinyID string `json:"account_destiny_id" example:"3"`
	Amount           int    `json:"amount" example:"1000"`
}
