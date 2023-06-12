package masta

import (
	"encoding/json"
	"strconv"
)

type ID = string

type Sbool bool

func (s *Sbool) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' && data[len(data)-1] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		b, err := strconv.ParseBool(str)
		if err != nil {
			return err
		}
		*s = Sbool(b)
		return nil
	}
	var b bool
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	*s = Sbool(b)
	return nil
}
