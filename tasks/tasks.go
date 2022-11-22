package tasks

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"pipeline/models"

	"github.com/sjsafranek/logger"
)

type ITask interface {
	Do(ctx context.Context) error
}

func New(ctx context.Context, method string, params *models.Params) (ITask, error) {
	return func() (ITask, error) {

		switch strings.ToLower(method) {

		case "etl":
			return &EtlTask{Params: params}, nil

		case "sleep":
			return &CallbackTask{Callback: func(ctx context.Context) error {
				time.Sleep(time.Duration(params.Timeout) * time.Millisecond)
				return nil
			}}, nil

		case "http_request":
			return &HttpRequestTask{Params: params}, nil

		case "create_directory", "create_directories", "mkdir", "mkdirs":
			if "" != params.Directory {
				params.Directories = append(params.Directories, params.Directory)
			}
			return &CallbackTask{Callback: func(ctx context.Context) error {
				for _, directory := range params.Directories {
					err := os.MkdirAll(directory, os.ModePerm)
					if nil != err {
						return err
					}
				}
				return nil
			}}, nil

		case "remove", "rm":
			// TODO patterm
			if "" != params.Directory {
				params.Directories = append(params.Directories, params.Directory)
			}
			if "" != params.File {
				params.Files = append(params.Files, params.File)
			}
			if "" != params.Filename {
				params.Files = append(params.Files, params.Filename)
			}
			items := append(params.Files, params.Directories...)
			return &CallbackTask{Callback: func(ctx context.Context) error {
				for _, item := range items {
					logger.Debugf("removing %v", item)
					err := os.Remove(item)
					if nil != err {
						return err
					}
				}
				return nil
			}}, nil

		case "move_file", "move_files", "mv":
			// TODO patterm
			return &CallbackTask{Callback: func(ctx context.Context) error {
				return os.Rename(params.InputFile, params.OutputFile)
			}}, nil

		case "copy_file", "cp":

		case "sql_query":

		default:
			return nil, fmt.Errorf("Unsuppored method: %v", method)

		}

		return nil, fmt.Errorf("Unsuppored method: %v", method)
	}()
}
