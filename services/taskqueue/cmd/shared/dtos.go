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
		Id          int        `json:"id"`
		UserId      int        `json:"userId"`
		PhoneNumber string     `json:"phoneNumber"`
		SmsBody     string     `json:"smsBody"`
		TryCount    int        `json:"tryCount"`
		Status      string     `json:"status"`
		CreatedDate *time.Time `json:"createdDate"`
	}
)
