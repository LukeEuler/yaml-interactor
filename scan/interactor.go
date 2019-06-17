package scan

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/LukeEuler/yaml-interactor/utils"
)

func interact(name, key string, context []string) yaml.MapItem {
	if len(context) > 0 {
		tip := strings.Join(context, ", ")
		println(utils.Green(tip))
	}
	buf := bufio.NewReader(os.Stdin)
	print(utils.Blue(name + ": "))
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	if len(sentence) <= 1 {
		println(utils.Red("Write something please."))
		return interact(name, key, context)
	}
	result := string(sentence[:len(sentence)-1])
	if num, err := strconv.Atoi(result); err == nil {
		return yaml.MapItem{Key: key, Value: num}
	} else {
		return yaml.MapItem{Key: key, Value: result}
	}
}
