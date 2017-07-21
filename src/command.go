package src

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "alert-conf-generator",
	Short: "alter configuration",
	Long:  `generator configure file for alter system in command line`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("source", "s", "", "configuration file to load")
	rootCmd.PersistentFlags().StringP("target", "t", "", "configuration file to write")
}

func Execute() error {
	return rootCmd.Execute()
}
