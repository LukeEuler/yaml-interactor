package src

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "yaml-file-generator",
	Short: "interact with yaml file",
	Long:  `complete a yaml file in command line`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("source", "s", "", "configure file to load")
	rootCmd.PersistentFlags().StringP("target", "t", "", "configure file to write")
}

func Execute() error {
	return rootCmd.Execute()
}
