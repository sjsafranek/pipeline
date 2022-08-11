package writers

import (
	"pipeline/models"
)

type MultiWriter struct {
	options *models.Options
	writers []IWriter
}

func (self *MultiWriter) Add(writer IWriter) {
	self.writers = append(self.writers, writer)
}

func (self *MultiWriter) WriteLine(row map[string]interface{}) {
	for _, writer := range self.writers {
		writer.WriteLine(row)
	}
}

func (self *MultiWriter) Close() {
	for _, writer := range self.writers {
		writer.Close()
	}
}
