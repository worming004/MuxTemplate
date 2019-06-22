package main

import (
	"encoding/json"
	"os"
)

// improve this configuration based on the new features you implement.
type configuration struct {
	allowedCors []string
}

func getConfiguration() configuration {
	confFile, err := os.Open("serverconfiguration.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	decoder := json.NewDecoder(confFile)
	config := configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}
