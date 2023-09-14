package service

import (
	"context"
	"fmt"
	"time"

	c "mid/core/var/common"
	"mid/services/taskqueue/cmd/shared"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) logMetrics(methodName string, err error, start time.Time) {

	lvs := []string{c.LabelMethod, methodName, c.LabelError, fmt.Sprint(err != nil)}
	mw.RequestCount.With(lvs...).Add(1)
	mw.RequestLatency.With(lvs...).Observe(time.Since(start).Seconds())

}

func (mw InstrumentingMiddleware) EnqueueSms(ctx context.Context, req shared.EnqueueSmsRequest) (resp shared.EnqueueSmsResponse, err error) {

	start := time.Now()
	defer mw.logMetrics(shared.EnqueueSmsMethodName, err, start)
	resp, err = mw.Next.EnqueueSms(ctx, req)
	return resp, err

}

func (mw InstrumentingMiddleware) TriggerWorker(ctx context.Context, req shared.TriggerWorkerRequest) (resp shared.TriggerWorkerResponse, err error) {

	start := time.Now()
	defer mw.logMetrics(shared.TriggerWorkerMethodName, err, start)
	resp, err = mw.Next.TriggerWorker(ctx, req)
	return resp, err

}

func (mw InstrumentingMiddleware) ReadAllSmsQueueFailed(ctx context.Context, req shared.ReadAllSmsQueueFailedRequest) (resp []shared.SmsQueue, err error) {

	start := time.Now()
	defer mw.logMetrics(shared.ReadAllSmsQueueFailedMethodName, err, start)
	resp, err = mw.Next.ReadAllSmsQueueFailed(ctx, req)
	return resp, err

}

func (mw InstrumentingMiddleware) ReadAllSmsQueue(ctx context.Context, req shared.ReadAllSmsQueueRequest) (resp []shared.SmsQueue, err error) {

	start := time.Now()
	defer mw.logMetrics(shared.ReadAllSmsQueueMethodName, err, start)
	resp, err = mw.Next.ReadAllSmsQueue(ctx, req)
	return resp, err

}
