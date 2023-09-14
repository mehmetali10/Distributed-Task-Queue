package service

import (
	"context"
	"fmt"
	"math/rand"
	"mid/core/util/loghelper"
	"mid/core/util/userhelper"
	"mid/services/taskqueue/cmd/shared"
	"mid/services/taskqueue/database/postgres"
	"mid/services/taskqueue/database/postgres/models/table"
	"sync"
	"time"

	"gorm.io/gorm"
)

type service struct{}

type Service interface {
	EnqueueSms(ctx context.Context, req shared.EnqueueSmsRequest) (shared.EnqueueSmsResponse, error)
	ReadAllSmsQueue(ctx context.Context, req shared.ReadAllSmsQueueRequest) ([]shared.SmsQueue, error)
	ReadAllSmsQueueFailed(ctx context.Context, req shared.ReadAllSmsQueueFailedRequest) ([]shared.SmsQueue, error)
	TriggerWorker(ctx context.Context, req shared.TriggerWorkerRequest) (shared.TriggerWorkerResponse, error)
}

func NewService() Service {
	return &service{}
}

func (*service) EnqueueSms(ctx context.Context, req shared.EnqueueSmsRequest) (shared.EnqueueSmsResponse, error) {

	logger := loghelper.NewLogger()
	currentUser, err := userhelper.ExctractUserFromContext(ctx)
	if err != nil {
		logger.Error("Error", "User not found")
		return shared.EnqueueSmsResponse{}, err
	}

	req.UserId = currentUser.UserID

	response, err := postgres.Create[shared.EnqueueSmsResponse, table.SmsQueue](ctx, req)
	if err != nil {
		logger.Error("Error", "Created Failed")
		return shared.EnqueueSmsResponse{}, err
	}
	logger.Info("Info", "A Sms Queued successfully")
	return response, nil

}

func (*service) ReadAllSmsQueue(ctx context.Context, req shared.ReadAllSmsQueueRequest) ([]shared.SmsQueue, error) {

	logger := loghelper.NewLogger()
	collection, err := postgres.Read[[]shared.SmsQueue, table.SmsQueue](ctx, map[string]interface{}{})
	if err != nil {
		logger.Error("Error", "Read Failed")
		return []shared.SmsQueue{}, err
	}
	logger.Info("Info", "All Queue Retrieved")
	return collection, nil

}

func (*service) ReadAllSmsQueueFailed(ctx context.Context, req shared.ReadAllSmsQueueFailedRequest) ([]shared.SmsQueue, error) {

	logger := loghelper.NewLogger()
	collection, err := postgres.Read[[]shared.SmsQueue, table.SmsQueue](ctx, map[string]interface{}{"Status": table.StatusFailed})
	if err != nil {
		logger.Error("Error", "Read Failed")
		return []shared.SmsQueue{}, err
	}
	logger.Info("Info", "All the failed Sms Queue Retrieved")
	return collection, nil

}

func (*service) TriggerWorker(ctx context.Context, req shared.TriggerWorkerRequest) (shared.TriggerWorkerResponse, error) {

	logger := loghelper.NewLogger()
	duration := 20
	wait := 5
	startTime := time.Now()
	handledSmsCount := 0

	for time.Since(startTime) < time.Duration(duration)*time.Second {

		PlayWorker(ctx, req, &handledSmsCount)
		logger.Info("Info", "PlayWorker completed successfully")

		if time.Since(startTime) > time.Duration(duration)*time.Second {
			logger.Info("Info", "Time is up")
			break
		}

		logger.Info("Info", "Sleep started for 5 second")
		time.Sleep(time.Duration(time.Second * time.Duration(wait)))
		logger.Info("Info", "Sleep ended for 5 second")

		if time.Since(startTime) > time.Duration(duration)*time.Second {
			logger.Info("Info", "Time is up")
			break
		}

	}

	return shared.TriggerWorkerResponse{HandledSmsCount: handledSmsCount}, nil

}

// PlayWorker orchestrates the execution of worker tasks for processing SMS queue items.
// It retrieves SMS queue entries using the GetQueue function, then spawns worker goroutines
// to process each queue item concurrently. The number of successfully processed items is
// subtracted from the 'count' parameter, and any errors encountered during processing are
// counted. The final count is updated in the 'count' parameter. If an error occurs during
// queue retrieval, it returns the error.
func PlayWorker(ctx context.Context, req shared.TriggerWorkerRequest, count *int) error {

	queue, err := GetQueue(ctx)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	resultChan := make(chan error, len(queue))

	for _, item := range queue {
		wg.Add(1)
		go Worker(ctx, item, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	errCount := 0
	for err := range resultChan {
		if err != nil {
			errCount++
		}
	}

	*count += len(queue) - errCount

	return nil

}

// Worker is a function used to process an SMS queue item.
// The function processes the given SMS queue item, sends an SMS (currently simulating a potential error),
// and reports the result. It signals the WaitGroup when the task is completed and conveys the results through a channel.
// If the item is successfully processed, its status is set to "Success"; otherwise, it is set to "Failed."
// Additionally, it increments the retry count and performs an item update in the database.
func Worker(ctx context.Context, item table.SmsQueue, wg *sync.WaitGroup, resultChan chan<- error) {

	defer wg.Done()
	// TODO: Send SMS here
	// For now, simulate a potential error
	err := GenerateRandomErrorWithProbability()
	if err != nil {
		item.Status = table.StatusFailed
		resultChan <- err
	} else {
		item.Status = table.StatusSuccess
	}

	item.TryCount++

	_, err = postgres.Update[table.SmsQueue, table.SmsQueue](ctx, map[string]interface{}{"Id": item.Id}, item)
	if err != nil {
		resultChan <- err
	}

}

// GetQueue retrieves SMS queue entries from the database based on certain criteria.
// It returns a slice of SmsQueue objects that have a "Status" of "PENDING" or "FAILED"
// and a "TryCount" less than 5. If an error occurs during the database query,
// it returns an empty slice and the error.
func GetQueue(ctx context.Context) ([]table.SmsQueue, error) {

	queue, err := postgres.Read[[]table.SmsQueue, table.SmsQueue](ctx, gorm.Expr("(\"Status\" = ? OR \"Status\" = ?) AND \"TryCount\" < ?", "PENDING", "FAILED", 5))

	if err != nil {
		return []table.SmsQueue{}, err
	}

	return queue, nil

}

// GenerateRandomErrorWithProbability generates a random error with a certain probability.
// It uses a random number generator to generate a number between 0 and 99 (inclusive).
// If the generated number is less than 50, it returns an error message indicating that a random error occurred.
// Otherwise, it returns nil, indicating no error.
func GenerateRandomErrorWithProbability() error {

	logger := loghelper.NewLogger()
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100)

	if randomNumber < 50 {
		logger.Info("Info", "Random error occurred")
		return fmt.Errorf("Random error occurred")
	}

	logger.Info("Info", "Random error null")
	return nil

}
