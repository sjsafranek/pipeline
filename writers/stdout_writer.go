package writers

import (
	"fmt"

	"pipeline/models"
)

type StdOutWriter struct {
	options *models.Options
}

func (self *StdOutWriter) WriteLine(row map[string]interface{}) error {
	fmt.Printf("%v\n", row)
	return nil
}

func (self *StdOutWriter) Close() {}
