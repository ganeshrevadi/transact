package payment

type status string


const (
		StatusCreated    Status = "CREATED"
    StatusProcessing Status = "PROCESSING"
    StatusSucceeded  Status = "SUCCEEDED"
    StatusFailed     Status = "FAILED"
)

type Payment struct(
	ID string `json:"id"`
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	Status Status `json:"status"`
)
