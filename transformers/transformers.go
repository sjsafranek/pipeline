package transformers

import (
	"fmt"

	"pipeline/models"
	"pipeline/readers"
)

type ITransformer interface {
	readers.IReader
	Transform(map[string]interface{}) map[string]interface{}
	Close()
}

func New(name string, reader readers.IReader, options *models.Options) (readers.IReader, error) {
	switch name {
	case "FilterTransformer":
		return &FilterTransformer{reader: reader, options: options}, nil
	case "AddColumnsTransformer":
		return &AddColumnsTransformer{reader: reader, options: options}, nil
	case "RemoveColumnsTransformer":
		return &RemoveColumnsTransformer{reader: reader, options: options}, nil
	case "AlterColumnsTransformer":
		return &AlterColumnsTransformer{reader: reader, options: options}, nil
	case "DataTypeTransformer":
	default:
		return nil, fmt.Errorf("Unsuppored reader: %v", name)
	}
	return nil, fmt.Errorf("Unsuppored reader: %v", name)
}
