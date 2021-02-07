package utils

import "github.com/goccy/go-json"

func DeepCopy(dst interface{}, src interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return err
	}
	return nil
}
