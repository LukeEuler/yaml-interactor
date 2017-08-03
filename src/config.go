package src

import (
	"yaml-interactor/scan"
	"yaml-interactor/utils"
)

func initConfig() {
	targetFile := rootCmd.PersistentFlags().Lookup("target").Value.String()
	utils.Writable(targetFile)

	sourceFile := rootCmd.PersistentFlags().Lookup("source").Value.String()
	out := scan.Handle(sourceFile)

	utils.WriteFile(targetFile, out)
}
