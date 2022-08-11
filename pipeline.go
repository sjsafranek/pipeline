package main

import (
	"context"

	"pipeline/models"
	"pipeline/tasks"
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
		for _, task := range self.Tasks {
			err := task.Do(ctx)
			if nil != err {
				return err
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
