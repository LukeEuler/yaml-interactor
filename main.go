package main

import (
	"flag"

	"github.com/LukeEuler/yaml-interactor/scan"
	"github.com/LukeEuler/yaml-interactor/utils"
)

func main() {
	sourceFile := flag.String("s", "source.yaml", "configure file to load")
	targetFile := flag.String("t", "target.yaml", "configure file to write")
	flag.Parse()

	utils.Writable(*targetFile)

	out := scan.Handle(*sourceFile)

	utils.WriteFile(*targetFile, out)
}
