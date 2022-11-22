package models

type Params struct {
	Readers      []*Options `json:"readers"`
	Transformers []*Options `json:"transformers"`
	Writers      []*Options `json:"writers"`
	Timeout      int64      `json:"timeout"`
	OutputFile   string     `json:"output_file"`
	Url          string     `json:"url"`
	Method       string     `json:"method"`
	Filename     string     `json:"filename"`
	Parallelize  bool       `json:"parallelize"`
	Directory    string     `json:"directory`
}
