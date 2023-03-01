package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Subject struct {
	Title    string `json:"title"`
	Symptom  string `json:"symptom"`
	Positive bool   `json:"positive"`
}

type Conf struct {
	Subject []Subject
	Options []string
}

var c Conf

func InitConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		panic(err)
	}
}
