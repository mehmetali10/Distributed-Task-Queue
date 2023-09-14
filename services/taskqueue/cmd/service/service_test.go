package service_test

import (
	"context"
	"mid/core/auth"
	"mid/core/var/common"
	"mid/services/taskqueue/cmd/service"
	"mid/services/taskqueue/cmd/shared"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/assert"
)

var svc = getService()
var ctx = getContext()

func TestEnqueue(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		resp, err := svc.EnqueueSms(ctx, shared.EnqueueSmsRequest{
			PhoneNumber: user.Phone,
			SmsBody:     "service_test",
		})
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})

	t.Run("Bad", func(t *testing.T) {
		ctx2 := context.Background()
		ctx2 = context.WithValue(ctx2, "user", &auth.User{})
		_, err := svc.EnqueueSms(ctx2, shared.EnqueueSmsRequest{
			PhoneNumber: user.Phone,
			SmsBody:     "service_test",
		})
		assert.NotNil(t, err)
	})

}

func TestReadAllFailSmsQueue(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		resp, err := svc.ReadAllSmsQueueFailed(ctx, shared.ReadAllSmsQueueFailedRequest{})
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})

}

func TestReadAllSmsQueue(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		resp, err := svc.ReadAllSmsQueue(ctx, shared.ReadAllSmsQueueRequest{})
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})

}

func TestTriggerWorker(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		resp, err := svc.TriggerWorker(ctx, shared.TriggerWorkerRequest{})
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})

}

func getService() service.Service {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "TimeStamp", log.DefaultTimestampUTC)

	fieldKeys := []string{common.LabelMethod, common.LabelError}
	requestCount := prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: shared.NameSpaceTaskQueueMicroservice,
		Name:      common.RequestCountName,
		Help:      common.RequestCountHelp,
	}, fieldKeys)

	requestLatency := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: shared.NameSpaceTaskQueueMicroservice,
		Name:      common.RequestLatencyName,
		Help:      common.RequestLatencyHelp,
	}, fieldKeys)

	countResult := prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: shared.NameSpaceTaskQueueMicroservice,
		Name:      common.CountResultName,
		Help:      common.CountResultHelp,
	}, fieldKeys)

	s := service.NewService()
	s = &service.LoggingMiddleware{logger, s}
	s = &service.InstrumentingMiddleware{requestCount, requestLatency, countResult, s}

	return s
}

var user = &auth.User{
	UserID:         99,
	Name:           "John",
	Surname:        "Doe",
	Phone:          "+905555555555",
	Email:          "johnDoe@test.com",
	StandardClaims: jwt.StandardClaims{},
}

func getContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user", user)
	return ctx
}
