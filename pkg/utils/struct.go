package utils

import "github.com/goccy/go-json"

// DeepCopy is a function that completely copies (without passing by reference) even the elements contained in a structure.
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
