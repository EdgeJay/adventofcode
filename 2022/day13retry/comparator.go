package day13retry

import (
	"encoding/json"
)

type JsonArray []interface{}

type Comparator struct {
	List []interface{}
}

func (c *Comparator) Parse(str string) error {

	list := JsonArray{}
	err := json.Unmarshal([]byte(str), &list)
	if err != nil {
		return err
	}

	c.List = list

	return nil
}
