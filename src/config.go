package src

import (
	"log"

	"gopkg.in/yaml.v2"
)

func initConfig() {
	targetFile := rootCmd.PersistentFlags().Lookup("target").Value.String()
	writableCheck(targetFile)

	sourceFile := rootCmd.PersistentFlags().Lookup("source").Value.String()
	sourceContent := readFile(sourceFile)

	data := unmarshal(sourceContent)
	data = scan(make([]string, 0), make([]string, 0), data)
	out := marshal(data)

	writeFile(targetFile, out)
}

func unmarshal(content []byte) yaml.MapSlice {
	data := yaml.MapSlice{}
	err := yaml.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func marshal(data yaml.MapSlice) []byte {
	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
