package models

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ParseFloat64(value interface{}) float64 {
	switch i := value.(type) {
	case float64:
		return i
	case float32:
		return float64(i)
	case int64:
		return float64(i)
	// ...other cases...
	default:
		return math.NaN()
	}
}

type Filter struct {
	Logical    string    `json:"logical"`
	Test       string    `json:"test"`
	Conditions []Filter  `json:"conditions`
	ColumnId   string    `json:"column_id"`
	Values     []string  `json:"values"`
	Value      string    `json:"value"`
	Min        float64   `json:"min"`
	Max        float64   `json:"max"`
	Begin      time.Time `json:"begin"`
	End        time.Time `json:"end"`
}

func (self *Filter) IsTest() bool {
	return "" != self.Test && "" == self.Logical
}

func (self *Filter) IsLogicalBlock() bool {
	return "" == self.Test && "" != self.Logical
}

func (self *Filter) Check(row map[string]interface{}) bool {
	if self.IsLogicalBlock() {
		switch strings.ToLower(self.Logical) {
		case "and":
			for _, filter := range self.Conditions {
				if !filter.Check(row) {
					return false
				}
			}
			return true
		case "or":
			for _, filter := range self.Conditions {
				if filter.Check(row) {
					return true
				}
			}
			return false
		case "not":
			// todo
		default:
			panic(fmt.Errorf("Unsuppored logical block: %v", self.Logical))
		}
	} else if self.IsTest() {
		switch strings.ToLower(self.Test) {
		case "in":
			return Contains(self.Values, fmt.Sprintf("%v", row[self.ColumnId]))
		case "not_in":
			return !Contains(self.Values, fmt.Sprintf("%v", row[self.ColumnId]))
		case "equals":
			return self.Value == row[self.ColumnId]
		case "not_equals":
			return self.Value == row[self.ColumnId]
		case "is_null":
			return nil == row[self.ColumnId] || "" == row[self.ColumnId]
		case "not_null":
			return nil != row[self.ColumnId] || "" == row[self.ColumnId]
		case "range":
			value := ParseFloat64(row[self.ColumnId])
			return (self.Min <= value) && (self.Max >= value)
		case "date_range":
			// TODO
		default:
			panic(fmt.Errorf("Unsuppored test: %v", self.Test))
		}
	}
	return true
}
