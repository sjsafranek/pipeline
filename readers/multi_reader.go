package readers

import (
	"pipeline/models"
)

type MultiReader struct {
	options *models.Options
	readers []IReader
}

func (self *MultiReader) Add(reader IReader) {
	self.readers = append(self.readers, reader)
}

func (self *MultiReader) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {
		for _, reader := range self.readers {
			for row := range reader.ReadLines() {
				queue <- row
			}
		}
		close(queue)
	}()
	return queue
}

func (self *MultiReader) Close() {
	for _, reader := range self.readers {
		reader.Close()
	}
}
