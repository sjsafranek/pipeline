package tasks

import (
	"context"
)

type CallbackTask struct {
	Callback func(ctx context.Context) error
}

func (self *CallbackTask) Do(ctx context.Context) error {
	return self.Callback(ctx)
}
