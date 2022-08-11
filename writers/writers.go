package writers

import (
	"fmt"

	"pipeline/models"
)

type IWriter interface {
	WriteLine(map[string]interface{})
	Close()
}

func New(name string, options *models.Options) (IWriter, error) {
	switch name {
	case "StdOutWriter":
		return &StdOutWriter{options: options}, nil
	case "CsvWriter":
		return &CsvWriter{options: options}, nil
	case "JsonWriter":
		return &JsonWriter{options: options}, nil
	// case "MultiWriter":
	// 	return &MultiWriter{options: options}, nil
	default:
		return nil, fmt.Errorf("Unsuppored writer: %v", name)
	}
}
