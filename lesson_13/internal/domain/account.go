package domain

type Account struct {
	ID      int64   `json:"id"`
	Balance float64 `json:"balance"`
	OwnerId int64   `json:"owner_id"`
}