package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"pipeline/models"
	"pipeline/tasks"

	"github.com/sjsafranek/logger"
)

type Pipeline struct {
	Name   string         `json:"name"`
	Method string         `json:"method`
	Tasks  []Pipeline     `json:"tasks"`
	Params *models.Params `json:"params"`
}

func (self *Pipeline) IsTask() bool {
	return "" != self.Method
}

func (self *Pipeline) Do(ctx context.Context) error {

	if !self.IsTask() {

		if nil != self.Params && self.Params.Parallelize {

			// Run tasks in parallel
			var wg = sync.WaitGroup{}
			errs := make(chan error, len(self.Tasks))
			for _, task := range self.Tasks {
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := task.Do(ctx)
					if nil != err {
						errs <- err
					}
				}()
			}
			wg.Wait()
			close(errs)

			// TODO: Bundle errors into a single error
			var err1 error
			for err2 := range errs {
				if nil == err1 {
					err1 = errors.New("Errors occured while processing task list")
				}
				err1 = fmt.Errorf("%v; %v", err1, err2)
			}
			return err1

		} else {

			// Run tasks sequantially
			for _, task := range self.Tasks {
				err := task.Do(ctx)
				if nil != err {
					logger.Errorf("%s: %s", task.Method, err.Error())
					return err
				}
			}

		}

		return nil
	}

	task, err := tasks.New(ctx, self.Method, self.Params)
	if nil != err {
		return err
	}

	return task.Do(ctx)
}
