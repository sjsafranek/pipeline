package tasks

type CallbackTask struct {
	Callback func() error
}

func (self *CallbackTask) Do() error {
	return self.Callback()
}
