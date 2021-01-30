package redmine

import (
	"encoding/json"
	"fmt"
	"github.com/goccy/go-json"
)

var dat map[string]interface{}

func Get(data []byte) {
	err := json.Unmarshal(data, dat)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)
}
