package transformers

import (
	"math"
	"strings"

	"pipeline/models"
	"pipeline/readers"
)

type AlterColumnsTransformer struct {
	reader  readers.IReader
	options *models.Options
}

func (self *AlterColumnsTransformer) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {
		for row := range self.reader.ReadLines() {
			queue <- self.Transform(row)
		}
		close(queue)
	}()
	return queue
}

func (self *AlterColumnsTransformer) Transform(row map[string]interface{}) map[string]interface{} {
	for _, column := range self.options.GetColumns() {
		switch strings.ToLower(column.Type) {
		case "float64":
			row[column.ColumnId] = self.ToFloat64(row[column.ColumnId])
		case "int64":
		default:
		}
	}
	return row
}

func (self AlterColumnsTransformer) ToFloat64(value interface{}) float64 {
	switch value.(type) {
	case int:
		return float64(value.(int))
	case float64:
		return value.(float64)
	default:
		return math.NaN()
	}
}

func (self *AlterColumnsTransformer) Close() {
	self.reader.Close()
}
