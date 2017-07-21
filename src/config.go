package src

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func initConfig() {
	target := rootCmd.PersistentFlags().Lookup("target").Value.String()
	targetCheck(target)

	source := rootCmd.PersistentFlags().Lookup("source").Value.String()
	sourceFile := sourceCheck(source)

	data := yaml.MapSlice{}
	err := yaml.Unmarshal(sourceFile, &data)
	if err != nil {
		log.Fatal(err)
	}
	data = scan(make([]string, 0), make([]string, 0), data)
	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(target, out, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

