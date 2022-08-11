package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (self *Pipeline) FromJSON(data []byte) error {
	return json.Unmarshal(data, self)
}

func (self *Pipeline) ToJSON() ([]byte, error) {
	return json.Marshal(self)
}

func (self *Pipeline) Load(filename string) error {
	file, err := os.Open(filename)
	if nil != err {
		return err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if nil != err {
		return err
	}
	return self.FromJSON(b)
}
