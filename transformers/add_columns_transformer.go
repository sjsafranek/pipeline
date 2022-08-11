package transformers

import (
	"pipeline/models"
	"pipeline/readers"
)

type AddColumnsTransformer struct {
	reader  readers.IReader
	options *models.Options
}

func (self *AddColumnsTransformer) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {
		for row := range self.reader.ReadLines() {
			queue <- self.Transform(row)
		}
		close(queue)
	}()
	return queue
}

func (self *AddColumnsTransformer) Transform(row map[string]interface{}) map[string]interface{} {
	for _, column := range self.options.GetColumns() {
		row[column.ColumnId] = column.Default
	}
	return row
}

func (self *AddColumnsTransformer) Close() {
	self.reader.Close()
}
