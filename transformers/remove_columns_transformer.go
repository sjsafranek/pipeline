package transformers

import (
	"pipeline/models"
	"pipeline/readers"
)

type RemoveColumnsTransformer struct {
	reader  readers.IReader
	options *models.Options
}

func (self *RemoveColumnsTransformer) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {
		for row := range self.reader.ReadLines() {
			queue <- self.Transform(row)
		}
		close(queue)
	}()
	return queue
}

func (self *RemoveColumnsTransformer) Transform(row map[string]interface{}) map[string]interface{} {
	for _, column := range self.options.GetColumns() {
		delete(row, column.ColumnId)
	}
	return row
}

func (self *RemoveColumnsTransformer) Close() {
	self.reader.Close()
}
