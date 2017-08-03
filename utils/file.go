package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

// Writable XXX
func Writable(target string) {
	_, err := os.Stat(target)
	if err != nil {
		_, err = os.Create(target)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = syscall.Access(target, syscall.O_RDWR)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s access ok\n", target)
	}
}

// ReadFile XXX
func ReadFile(source string) []byte {
	sourceFile, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	return sourceFile
}

// WriteFile XXX
func WriteFile(file string, content []byte) {
	err := ioutil.WriteFile(file, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
