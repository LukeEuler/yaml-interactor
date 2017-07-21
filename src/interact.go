package src

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func interact(name string, context []string) string {
	if len(context) > 0 {
		tip := strings.Join(context, ", ")
		println(green(tip))
	}
	buf := bufio.NewReader(os.Stdin)
	print(blue(name + ": "))
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	if len(sentence) <= 1 {
		println(red("Write something please."))
		return interact(name, context)
	}
	result := string(sentence[:len(sentence)-1])
	//result = strings.TrimSuffix(result, "\n")
	return result
}
