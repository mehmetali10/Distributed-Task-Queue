package shared

import "time"

// EnqueueSms
type (
	EnqueueSmsRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		SmsBody     string `json:"smsBody" validate:"required"`
		UserId      int    `swaggerignore:"true"`
	}

	EnqueueSmsResponse struct {
		Id     int    `json:"id"`
		Status string `json:"status"`
	}
)

// TriggerWorker
type (
	TriggerWorkerRequest struct{}

	TriggerWorkerResponse struct {
		HandledSmsCount int `json:"handledSmsCount"`
	}
)

// Read SmsQueue
type (
	ReadAllSmsQueueRequest       struct{}
	ReadAllSmsQueueFailedRequest struct{}

	SmsQueue struct {
		Id          int
		UserId      int
		PhoneNumber string
		SmsBody     string
		TryCount    int
		Status      string
		CreatedDate *time.Time
	}
)
