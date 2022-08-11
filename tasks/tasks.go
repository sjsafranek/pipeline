package tasks

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"pipeline/models"
)

type ITask interface {
	Do(ctx context.Context) error
}

func New(ctx context.Context, method string, params *models.Params) (ITask, error) {
	switch strings.ToLower(method) {
	case "etl":
		return &EtlTask{Params: params}, nil
	case "sleep":
		return &CallbackTask{Callback: func(ctx context.Context) error {
			time.Sleep(time.Duration(params.Timeout) * time.Millisecond)
			return nil
		}}, nil
	case "http_request":
		return &HttpRequestTask{Params: params}, nil
	case "delete_file":
		return &CallbackTask{Callback: func(ctx context.Context) error {
			return os.Remove(params.Filename)
		}}, nil
	case "move_file":
	case "copy_file":
	case "sql_query":
	default:
		return nil, fmt.Errorf("Unsuppored method: %v", method)
	}
	return nil, fmt.Errorf("Unsuppored method: %v", method)
}
