package models

type Params struct {
	Readers      []*Options `json:"readers"`
	Transformers []*Options `json:"transformers"`
	Writers      []*Options `json:"writers"`

	Timeout int64 `json:"timeout"`

	Filename   string   `json:"filename"`
	File       string   `json:"file"`
	Files      []string `json:"files"`
	InputFile  string   `json:"input_file"`
	OutputFile string   `json:"output_file"`

	Url string `json:"url"`

	Method string `json:"method"`

	Parallelize bool `json:"parallelize"`

	Directory   string   `json:"directory`
	Directories []string `json:"directories"`
}
