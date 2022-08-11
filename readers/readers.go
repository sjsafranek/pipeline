package readers

import (
	"fmt"
	//"errors"

	"pipeline/models"
)

type IReader interface {
	ReadLines() chan map[string]interface{}
	Close()
}

func New(name string, options *models.Options) (IReader, error) {
	switch name {
	case "CsvReader":
		return &CsvReader{options: options}, nil
	default:
		return nil, fmt.Errorf("Unsuppored reader: %v", name)
	}
}
