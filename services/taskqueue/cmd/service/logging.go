package service

import (
	"context"
	c "mid/core/var/common"
	"mid/services/taskqueue/cmd/shared"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

func (mw *LoggingMiddleware) logRequestResponse(methodName string, err error, start time.Time) {

	_ = level.Info(mw.Logger).Log(
		c.LabelMethod, methodName,
		c.LabelErr, err,
		c.LabelTook, time.Since(start),
	)

}

func (mw *LoggingMiddleware) EnqueueSms(ctx context.Context, req shared.EnqueueSmsRequest) (resp shared.EnqueueSmsResponse, err error) {

	start := time.Now()
	defer mw.logRequestResponse(shared.EnqueueSmsMethodName, err, start)
	return mw.Next.EnqueueSms(ctx, req)

}

func (mw *LoggingMiddleware) TriggerWorker(ctx context.Context, req shared.TriggerWorkerRequest) (resp shared.TriggerWorkerResponse, err error) {

	start := time.Now()
	defer mw.logRequestResponse(shared.TriggerWorkerMethodName, err, start)
	return mw.Next.TriggerWorker(ctx, req)

}

func (mw *LoggingMiddleware) ReadAllSmsQueue(ctx context.Context, req shared.ReadAllSmsQueueRequest) (resp []shared.SmsQueue, err error) {

	start := time.Now()
	defer mw.logRequestResponse(shared.ReadAllSmsQueueMethodName, err, start)
	return mw.Next.ReadAllSmsQueue(ctx, req)

}

func (mw *LoggingMiddleware) ReadAllSmsQueueFailed(ctx context.Context, req shared.ReadAllSmsQueueFailedRequest) (resp []shared.SmsQueue, err error) {

	start := time.Now()
	defer mw.logRequestResponse(shared.ReadAllSmsQueueFailedMethodName, err, start)
	return mw.Next.ReadAllSmsQueueFailed(ctx, req)

}
