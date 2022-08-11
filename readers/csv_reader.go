package readers

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"pipeline/models"
)

type CsvReader struct {
	options *models.Options
}

func (self *CsvReader) ReadLines() chan map[string]interface{} {
	queue := make(chan map[string]interface{}, 1)
	go func() {

		// open file
		fh, err := os.Open(self.options.GetFilename())
		if err != nil {
			log.Fatal(err)
		}

		// remember to close the file at the end of the program
		defer fh.Close()

		// read csv values using csv.Reader
		header := []string{}
		reader := csv.NewReader(fh)
		reader.Comma = self.options.GetDelimiter()
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			if 0 == len(header) {
				header = line
				continue
			}

			// Build row
			row := make(map[string]interface{})
			for i, column_id := range header {
				row[column_id] = line[i]
			}

			queue <- row
		}

		close(queue)
	}()
	return queue
}

func (self *CsvReader) Close() {}
