package writers

import (
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"

	"pipeline/models"
)

type CsvWriter struct {
	options *models.Options
	fh      *os.File
	writer  *csv.Writer
	//stream *io.Writer
}

func (self *CsvWriter) WriteLine(row map[string]interface{}) {

	if nil == self.fh {
		fh, err := os.Create(self.options.GetFilename())
		if err != nil {
			log.Fatal(err)
		}
		self.fh = fh
		writer := csv.NewWriter(self.fh)
		self.writer = writer
		self.writer.Comma = self.options.GetDelimiter()
	}

	if 0 == len(self.options.Header) {
		for column_id := range row {
			self.options.Header = append(self.options.Header, column_id)
		}
		self.write(self.options.Header)
	}

	values := make([]string, len(self.options.Header))
	for i, column_id := range self.options.Header {
		values[i] = fmt.Sprintf("%v", row[column_id])
	}
	self.write(values)
}

func (self *CsvWriter) write(line []string) {
	err := self.writer.Write(line)
	if err != nil {
		log.Fatal(err)
	}
	self.writer.Flush()
}

func (self *CsvWriter) Close() {
	self.fh.Close()
}
