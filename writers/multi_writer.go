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

func (self *MultiWriter) WriteLine(row map[string]interface{}) error {
	for _, writer := range self.writers {
		err := writer.WriteLine(row)
		if nil != err {
			return err
		}
	}
	return nil
}

func (self *MultiWriter) Close() {
	for _, writer := range self.writers {
		writer.Close()
	}
}
