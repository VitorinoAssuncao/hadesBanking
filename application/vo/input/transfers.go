package input

type CreateTransferVO struct {
	AccountOrigin_ID  string `json:"account_origin_id" example:"2"`
	AccountDestiny_ID string `json:"account_destiny_id" example:"3"`
	Balance           int    `json:"balance" example:"1000"`
}
