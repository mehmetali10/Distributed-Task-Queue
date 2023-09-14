package endpoint_test

import (
	"context"
	"mid/core/auth"
	"mid/services/taskqueue/cmd/endpoint"
	"mid/services/taskqueue/cmd/service"
	"mid/services/taskqueue/cmd/shared"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var svc = getService()
var ctx = getContext()

func TestEndpoints(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		endpoint := endpoint.MakeEnqueueSmsEndpoint(svc)
		ctx2 := context.Background()
		ctx2 = context.WithValue(ctx2, "user", &auth.User{})
		endpoint(ctx2, shared.EnqueueSmsRequest{})
		assert.NotNil(t, endpoint)
	})

	t.Run("OK", func(t *testing.T) {
		endpoint := endpoint.MakeReadAllSmsQueueEndpoint(getService())
		endpoint(context.Background(), shared.ReadAllSmsQueueRequest{})
		assert.NotNil(t, endpoint)
	})

	t.Run("OK", func(t *testing.T) {
		endpoint := endpoint.MakeReadAllSmsQueueFailedEndpoint(getService())
		assert.NotNil(t, endpoint)
	})

	t.Run("OK", func(t *testing.T) {
		endpoint := endpoint.MakeTriggerWorkerEndpoint(getService())
		assert.NotNil(t, endpoint)
	})
}

func getService() service.Service {
	return service.NewService()
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
