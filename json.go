package github.com/tolgaerdonmez/goutils

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveJSON saves the given data to json
func SaveJSON(data interface{}, outputname string) {
	file, err := os.Create(outputname)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
