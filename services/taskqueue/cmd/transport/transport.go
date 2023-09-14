package transport

import (
	"context"
	"encoding/json"
	"mid/core/auth"
	customerrors "mid/core/errors"
	"mid/core/util/transporthelper"
	"mid/core/var/common"
	c "mid/core/var/common"
	e "mid/services/taskqueue/cmd/endpoint"
	s "mid/services/taskqueue/cmd/service"
	"mid/services/taskqueue/cmd/shared"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"golang.org/x/time/rate"
)

func authServer[T any](service s.Service, endpoint endpoint.Endpoint) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(10), 100)

	server := httptransport.NewServer(
		ratelimit.NewErroringLimiter(limiter)(endpoint),
		decodeRequest[T],
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	)

	handler := auth.AuthMiddleware(server)

	return handler
}

func EnqueueSmsServer(service s.Service) http.Handler {
	return authServer[shared.EnqueueSmsRequest](service, e.MakeEnqueueSmsEndpoint(service))
}

func TriggerWorkerServer(service s.Service) http.Handler {
	return authServer[shared.TriggerWorkerRequest](service, e.MakeTriggerWorkerEndpoint(service))
}

func ReadAllSmsQueueServer(service s.Service) http.Handler {
	return authServer[shared.ReadAllSmsQueueRequest](service, e.MakeReadAllSmsQueueEndpoint(service))
}

func ReadAllSmsQueueFailedServer(service s.Service) http.Handler {
	return authServer[shared.ReadAllSmsQueueFailedRequest](service, e.MakeReadAllSmsQueueFailedEndpoint(service))
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(common.LabelContentType, common.LabelApplicationJsonUtf8)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func decodeRequest[T any](ctx context.Context, r *http.Request) (interface{}, error) {
	var req T

	err1 := json.NewDecoder(r.Body).Decode(&req)
	if err1 != nil {
		if r.Body == http.NoBody {
			return req, nil
		}
		return nil, err1
	}

	switch any(req).(type) {
	case shared.EnqueueSmsRequest:
		rt, ok := any(req).(shared.EnqueueSmsRequest)
		if ok {
			condition := transporthelper.IsPhoneValid(rt.PhoneNumber)
			if !condition {
				return nil, &customerrors.InvalidPhoneNumber{}
			}
		}
	}

	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set(c.LabelContentType, c.LabelApplicationJsonUtf8)

	switch err.(type) {
	case validator.ValidationErrors:
		w.WriteHeader(http.StatusBadRequest)
	case *customerrors.InvalidPhoneNumber:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		c.LabelError: err.Error(),
	})
}
