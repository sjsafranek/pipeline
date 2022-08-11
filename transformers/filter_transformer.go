package transformers

import (
	"pipeline/models"
	"pipeline/readers"
)

type FilterTransformer struct {
	reader  readers.IReader
	options *models.Options
}

func (self *FilterTransformer) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {
		for row := range self.reader.ReadLines() {
			row = self.Transform(row)
			if nil != row {
				queue <- row
			}
		}
		close(queue)
	}()
	return queue
}

func (self *FilterTransformer) Transform(row map[string]interface{}) map[string]interface{} {
	filter := self.options.GetFilter()
	if filter.Check(row) {
		return row
	}
	return nil
}

func (self *FilterTransformer) Close() {
	self.reader.Close()
}
