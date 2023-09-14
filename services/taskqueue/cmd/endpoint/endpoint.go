package endpoint

import (
	"context"
	customerrors "mid/core/errors"
	s "mid/services/taskqueue/cmd/service"

	"github.com/go-kit/kit/endpoint"
)

func makeEndpoint[T any, K any](s s.Service, f func(ctx context.Context, req T) (resp K, err error)) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(T)
		if !ok {
			return nil, &customerrors.InvalidRequestType{}
		}
		response, err := f(ctx, req)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// EnqueueSms godoc
// @Tags         TaskQueue
// @Summary      Enqueue SMS
// @Description  This endpoint enqueues an SMS for processing in the task queue.
// @Accept       json
// @Produce      json
// @Param        req body shared.EnqueueSmsRequest true "Request object for enqueuing an SMS"
// @Security     BearerAuth
// @Success      200 {object} shared.EnqueueSmsResponse "Successful response with details of the enqueued SMS"
// @Failure      400
// @Router       /SmsQueue/Enqueue [post]
func MakeEnqueueSmsEndpoint(s s.Service) endpoint.Endpoint {
	return makeEndpoint(s, s.EnqueueSms)
}

// TriggerWorker godoc
// @Tags         TaskQueue
// @Summary      Trigger Worker
// @Description  This endpoint triggers a worker for processing tasks in the task queue.
// @Accept       json
// @Produce      json
// @Param        req body shared.TriggerWorkerRequest true "Request object for triggering a worker"
// @Security     BearerAuth
// @Success      200 {object} shared.TriggerWorkerResponse "Successful response with details of worker execution"
// @Failure      400
// @Router       /SmsQueue/TriggerWorker [post]
func MakeTriggerWorkerEndpoint(s s.Service) endpoint.Endpoint {
	return makeEndpoint(s, s.TriggerWorker)
}

// ReadAllSmsQueue godoc
// @Tags         TaskQueue
// @Summary      Read All SMS Queue
// @Description  This endpoint retrieves all SMS queue entries.
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} shared.SmsQueue "List of SMS queue entries"
// @Failure      400
// @Router       /SmsQueue/ReadAll [get]
func MakeReadAllSmsQueueEndpoint(s s.Service) endpoint.Endpoint {
	return makeEndpoint(s, s.ReadAllSmsQueue)
}

// ReadAllSmsQueueFailed godoc
// @Tags         TaskQueue
// @Summary      Read All Failed SMS Queue Entries
// @Description  This endpoint retrieves all failed SMS queue entries.
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} shared.SmsQueue "List of failed SMS queue entries"
// @Failure      400
// @Router       /SmsQueue/ReadAll/Fail [get]
func MakeReadAllSmsQueueFailedEndpoint(s s.Service) endpoint.Endpoint {
	return makeEndpoint(s, s.ReadAllSmsQueueFailed)
}
