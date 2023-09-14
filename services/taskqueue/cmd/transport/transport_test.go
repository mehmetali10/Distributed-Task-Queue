package transport_test

import (
	"mid/core/util/transporthelper"
	"mid/services/taskqueue/cmd/service"
	"mid/services/taskqueue/cmd/shared"
	"mid/services/taskqueue/cmd/transport"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriggerWorker(t *testing.T) {

	t.Run("Ok", func(t *testing.T) {
		server := httptest.NewServer(transport.TriggerWorkerServer(service.NewService()))
		bodyParam := shared.TriggerWorkerRequest{}

		b := transporthelper.BaseTemp{
			Server:    server,
			Url:       server.URL + shared.TriggerWorkerApiPath,
			Want:      http.StatusOK,
			Method:    http.MethodPost,
			BodyParam: bodyParam,
		}

		resp := transporthelper.TestWithToken(t, b)
		assert.NotNil(t, resp["totalRow"])
	})
}

func TestUnAuth(t *testing.T) {
	server := httptest.NewServer(transport.ReadAllSmsQueueServer(service.NewService()))
	bodyParam := shared.ReadAllSmsQueueRequest{}

	b := transporthelper.BaseTemp{
		Server:    server,
		Url:       server.URL + shared.ReadAllSmsQueueApiPath,
		Want:      http.StatusUnauthorized,
		Method:    http.MethodGet,
		BodyParam: bodyParam,
	}

	resp := transporthelper.TestWithoutToken(t, b)

	if resp["error"] == "" {
		t.Errorf("error empty")
	}
}

func TestReadAllSmsQueueValid(t *testing.T) {
	server := httptest.NewServer(transport.ReadAllSmsQueueServer(service.NewService()))
	bodyParam := shared.ReadAllSmsQueueRequest{}

	b := transporthelper.BaseTemp{
		Server:    server,
		Url:       server.URL + shared.ReadAllSmsQueueApiPath,
		Want:      http.StatusOK,
		Method:    http.MethodGet,
		BodyParam: bodyParam,
	}

	transporthelper.TestWithToken(t, b)
}

func TestReadAllSmsQueueFailValid(t *testing.T) {
	server := httptest.NewServer(transport.ReadAllSmsQueueFailedServer(service.NewService()))
	bodyParam := shared.ReadAllSmsQueueFailedRequest{}

	b := transporthelper.BaseTemp{
		Server:    server,
		Url:       server.URL + shared.ReadAllSmsQueueFailedApiPath,
		Want:      http.StatusOK,
		Method:    http.MethodGet,
		BodyParam: bodyParam,
	}

	transporthelper.TestWithToken(t, b)
}

func TestEnqueue(t *testing.T) {

	t.Run("Ok", func(t *testing.T) {
		server := httptest.NewServer(transport.EnqueueSmsServer(service.NewService()))
		bodyParam := shared.EnqueueSmsRequest{
			PhoneNumber: "+905555555555",
			SmsBody:     "Transport test",
		}

		b := transporthelper.BaseTemp{
			Server:    server,
			Url:       server.URL + shared.EnqueueSmsApiPath,
			Want:      http.StatusOK,
			Method:    http.MethodPost,
			BodyParam: bodyParam,
		}
		resp := transporthelper.TestWithToken(t, b)
		if resp["id"] == "" || resp["id"] == nil {
			t.Errorf("id nil or empty")
		}
	})

	t.Run("Bad", func(t *testing.T) {
		server := httptest.NewServer(transport.EnqueueSmsServer(service.NewService()))
		bodyParams := []shared.EnqueueSmsRequest{
			shared.EnqueueSmsRequest{
				// PhoneNumber: "+905555555555",
				SmsBody: "Transport test",
			},
			shared.EnqueueSmsRequest{
				PhoneNumber: "+905555555555",
				// SmsBody:     "Transport test",
			},
		}

		for _, v := range bodyParams {
			b := transporthelper.BaseTemp{
				Server:    server,
				Url:       server.URL + shared.EnqueueSmsApiPath,
				Want:      http.StatusBadRequest,
				Method:    http.MethodPost,
				BodyParam: v,
			}

			resp := transporthelper.TestWithToken(t, b)

			if resp["error"] == "" {
				t.Errorf("error empty")
			}
		}
	})
}
