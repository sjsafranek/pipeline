package models

type Column struct {
	ColumnId string      `json:"column_id"`
	Type     string      `json:"type"`
	Default  interface{} `json:"default"`
}
