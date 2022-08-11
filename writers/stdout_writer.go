package writers

import (
	"fmt"

	"pipeline/models"
)

type StdOutWriter struct {
	options *models.Options
}

func (self *StdOutWriter) WriteLine(row map[string]interface{}) {
	fmt.Printf("%v\n", row)
}

func (self *StdOutWriter) Close() {}
