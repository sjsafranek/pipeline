package tasks

import (
	"io"
	"net/http"
	"os"
	"context"

	"pipeline/models"
)

type HttpRequestTask struct {
	Params *models.Params `json:"params"`
}

func (self *HttpRequestTask) Do(ctx context.Context) error {
	// Create the file
	out, err := os.Create(self.Params.OutputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	// TODO: check self.Params.Method

	// Get the data
	var resp *http.Response
	resp, err = http.Get(self.Params.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
