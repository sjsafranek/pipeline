package tasks

import (
	"context"
)

type CallbackTask struct {
	Callback func(ctx context.Context) error
}

func (self *CallbackTask) Do(ctx context.Context) error {
	if nil != ctx.Err() {
		return ctx.Err()
	}
	err := self.Callback(ctx)
	if nil != err {
		return err
	}
	return ctx.Err()
}
