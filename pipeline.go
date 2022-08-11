package main

import (
	"pipeline/models"
	"pipeline/tasks"
)

type Pipeline struct {
	Method string         `json:"method`
	Tasks  []Pipeline     `json:"tasks"`
	Params *models.Params `json:"params"`
}

func (self *Pipeline) IsTask() bool {
	return "" != self.Method
}

func (self *Pipeline) Do() error {
	if !self.IsTask() {
		for _, task := range self.Tasks {
			err := task.Do()
			if nil != err {
				return err
			}
		}
		return nil
	}

	task, err := tasks.New(self.Method, self.Params)
	if nil != err {
		return err
	}
	return task.Do()
}
