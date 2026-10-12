package payment

import "errors"

type Status string

const (
	StatusCreated    Status = "CREATED"
	StatusProcessing Status = "PROCESSING"
	StatusSucceeded  Status = "SUCCEEDED"
	StatusFailed     Status = "FAILED"
)

type Payment struct {
	ID       string `json:"id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Status   Status `json:"status"`
}

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (r CreatePaymentRequest) Validate() error {
	if r.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if r.Currency == "" {
		return errors.New("currency is required")
	}
	return nil
}
