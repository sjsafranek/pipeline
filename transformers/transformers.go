package transformers

import (
	"fmt"

	"pipeline/models"
	"pipeline/readers"
)

type ITransformer interface {
	readers.IReader
	//reader readers.IReader
	Transform(chan map[string]interface{}) chan map[string]interface{}
	Close()
}

func New(name string, reader readers.IReader, options *models.Options) (readers.IReader, error) {
	switch name {
	case "FilterTransformer":
		return &FilterTransformer{reader: reader, options: options}, nil
	default:
		return nil, fmt.Errorf("Unsuppored reader: %v", name)
	}
}

// AddColumnsTransformer
// RemoveColumnsTransformer
// DataTypeTransformer
