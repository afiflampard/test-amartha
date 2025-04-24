package helper

import (
	"encoding/json"
	"log"
)

func PrintJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}
	log.Println(string(jsonData))
}
