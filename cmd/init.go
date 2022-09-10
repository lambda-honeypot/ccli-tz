package cmd

import (
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise $HOME/.ccli-tz.yaml config file",
	Long: `Create a sample config file at the path $HOME/.ccli-tz.yaml.

This has a minimal set of fields needed to run the cardano-cli query leadership-schedule.`,
	Run: initialiseConfigFileCmd,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(initCmd)
}

func initialiseConfigFileCmd(_ *cobra.Command, _ []string) {
	config.InitialiseConfigFile(&config.FileConfigCreator{})
}
