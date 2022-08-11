package models

type Options struct {
	Type      string   `json:"type"`
	Filename  string   `json:"filename"`
	Delimiter *rune    `json:"delimiter"`
	Level     string   `json:"level"`
	Header    []string `json:"header"`
	Filter    *Filter  `json:"filter"`
	Options   *Options `json:"options`
}

func (self *Options) GetFilename() string {
	return self.Filename
}

func (self *Options) GetDelimiter() rune {
	if nil != self.Delimiter {
		return *self.Delimiter
	}
	return ','
}

func (self *Options) GetLevel() string {
	if "" == self.Level {
		return "DEBUG"
	}
	return self.Level
}

func (self *Options) GetFilter() *Filter {
	if nil != self.Filter {
		return self.Filter
	}
	return &Filter{}
}
