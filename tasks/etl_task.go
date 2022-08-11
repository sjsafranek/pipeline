package tasks

import (
	"pipeline/models"
	"pipeline/readers"
	"pipeline/transformers"
	"pipeline/writers"
)

type EtlTask struct {
	Params *models.Params `json:"params"`
}

func (self *EtlTask) newReader() (readers.IReader, error) {
	if 1 == len(self.Params.Readers) {
		options := self.Params.Readers[0]
		reader, err := readers.New(options.Type, options.Options)
		if nil != err {
			return nil, err
		}
		return self.addTransformers(reader)
	}
	reader := readers.MultiReader{}
	for _, options := range self.Params.Writers {
		r, err := readers.New(options.Type, options.Options)
		if nil != err {
			return nil, err
		}
		reader.Add(r)
	}
	return self.addTransformers(&reader)
}

func (self *EtlTask) addTransformers(reader readers.IReader) (readers.IReader, error) {
	if 0 == len(self.Params.Transformers) {
		return reader, nil
	}
	for _, options := range self.Params.Transformers {
		r, err := transformers.New(options.Type, reader, options.Options)
		if nil != err {
			return nil, err
		}
		reader = r
	}
	return reader, nil
}

func (self *EtlTask) newWriter() (writers.IWriter, error) {
	if 1 == len(self.Params.Writers) {
		options := self.Params.Writers[0]
		return writers.New(options.Type, options.Options)
	}
	writer := writers.MultiWriter{}
	for _, options := range self.Params.Writers {
		w, err := writers.New(options.Type, options.Options)
		if nil != err {
			return nil, err
		}
		writer.Add(w)
	}
	return &writer, nil
}

func (self *EtlTask) Do() error {
	reader, err := self.newReader()
	if nil != err {
		return err
	}
	defer reader.Close()

	writer, err := self.newWriter()
	if nil != err {
		return err
	}
	defer writer.Close()

	for row := range reader.ReadLines() {
		writer.WriteLine(row)
		if nil != err {
			return err
		}
	}

	return nil
}
