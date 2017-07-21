package src

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
)

func targetCheck(target string) {
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

func sourceCheck(source string) []byte {
	sourceFile, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatal(err)
	}
	return sourceFile
}

