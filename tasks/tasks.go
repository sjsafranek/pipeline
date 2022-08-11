package tasks

import (
	"fmt"
	"strings"
	"time"

	"pipeline/models"
)

type ITask interface {
	Do() error
}

func New(method string, params *models.Params) (ITask, error) {
	switch strings.ToLower(method) {
	case "etl":
		return &EtlTask{Params: params}, nil
	case "sleep":
		return &CallbackTask{Callback: func() error {
			time.Sleep(time.Duration(params.Timeout) * time.Millisecond)
			return nil
		}}, nil
	case "http_request":
		return &HttpRequestTask{Params: params}, nil
	case "sql_query":
	case "move_file":
	case "copy_file":
	default:
		return nil, fmt.Errorf("Unsuppored method: %v", method)
	}
	return nil, fmt.Errorf("Unsuppored method: %v", method)
}
