package writers

import (
	"encoding/json"
	"log"
	"os"
	//"io"

	"pipeline/models"
)

type JsonWriter struct {
	options *models.Options
	fh      *os.File
	//stream *io.Writer
}

func (self *JsonWriter) WriteLine(row map[string]interface{}) {

	first := false
	if nil == self.fh {
		fh, err := os.Create(self.options.GetFilename())
		if err != nil {
			log.Fatal(err)
		}
		self.fh = fh
		self.fh.Write([]byte("["))
		first = true
	}

	b, err := json.Marshal(row)
	if err != nil {
		log.Fatal(err)
	}

	if first {
		self.fh.Write([]byte("\n\t"))
		self.fh.Write(b)
	} else {

		self.fh.Write([]byte(",\n\t"))
		self.fh.Write(b)
	}

}

func (self *JsonWriter) Close() {
	self.fh.Write([]byte("\n]"))
	self.fh.Close()
}
