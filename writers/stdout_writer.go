package writers

import (
	"fmt"

	"pipeline/models"
)

type StdOutWriter struct {
	options *models.Options
}

func (self *StdOutWriter) WriteLine(row map[string]interface{}) error {
	// apply filter
	if nil != self.options && nil != self.options.Filter && !self.options.Filter.Check(row) {
		return nil
	}

	fmt.Printf("%v\n", row)
	return nil
}

func (self *StdOutWriter) Close() {}
