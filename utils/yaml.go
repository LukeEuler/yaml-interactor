package utils

import (
	"log"

	"gopkg.in/yaml.v2"
)

// Unmarshal XXX
func Unmarshal(content []byte) yaml.MapSlice {
	data := yaml.MapSlice{}
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Marshal XXX
func Marshal(data yaml.MapSlice) []byte {
	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
