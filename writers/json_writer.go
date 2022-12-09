package writers

import (
	"encoding/json"
	"os"
	//"io"

	"pipeline/models"
)

var json_start_prefix = []byte("[\n\t")
var json_middle_prefix = []byte(",\n\t")

type JsonWriter struct {
	options *models.Options
	fh      *os.File
	//stream *io.Writer
}

func (self *JsonWriter) WriteLine(row map[string]interface{}) error {
	// apply filter
	if nil != self.options &&  nil != self.options.Filter && !self.options.Filter.Check(row) {
		return nil
	}

	prefix := json_middle_prefix

	if nil == self.fh {
		fh, err := os.Create(self.options.GetFilename())
		if nil != err {
			return err
		}
		self.fh = fh
		prefix = json_start_prefix
	}

	data, err := json.Marshal(row)
	if nil != err {
		return err
	}

	_, err = self.fh.Write(append([]byte(prefix), data...))
	return err
}

func (self *JsonWriter) Close() {
	self.fh.Write([]byte("\n]"))
	self.fh.Close()
}
