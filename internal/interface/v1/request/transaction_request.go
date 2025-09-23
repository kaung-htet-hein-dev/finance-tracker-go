package request

type CreateTransactionRequest struct {
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	Note       *string `json:"note" validate:"omitempty,max=255"`
	Type       string  `json:"type" validate:"required,oneof=income expense"`
	CategoryID uint    `json:"category_id" validate:"required,gt=0"`
}

type UpdateTransactionRequest struct {
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	Note       string  `json:"note" validate:"omitempty,max=255"`
	Type       string  `json:"type" validate:"required,oneof=income expense"`
	CategoryID uint    `json:"category_id" validate:"required,gt=0"`
}

type FilterTransactionsRequest struct {
	Type       string `json:"type" validate:"omitempty,oneof=income expense"`
	CategoryID uint   `json:"category_id" validate:"omitempty,gte=0"`
}
