package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Data structure to hold the JSON data
type Data struct {
	Countries map[string]map[string][]string `json:""`
}

// LoadData function to read and parse the JSON file
func LoadData(filename string) Data {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to load data file: %v", err)
	}

	var data Data
	if err := json.Unmarshal(file, &data.Countries); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return data
}
