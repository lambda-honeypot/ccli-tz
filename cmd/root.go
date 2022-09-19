package cmd

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccli-tz",
	Short: "Wrapper around the cardano-cli leadership schedule with timezone conversion",
	Long: `Wraps around the cardano-cli query leadership-schedule and allows for timezone conversion of dates. For example:

ccli-tz --timezone europe/london --next

This will create the leadership-schedule for the pool and keys within the config file for the next epoch.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: LeadershipLog,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func LeadershipLog(_ *cobra.Command, _ []string) {
	timeZone := "America/New_York"
	leader.CreateAndRun(timeZone, &leader.CmdRunner{})
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccli-tz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().Bool("current", false, "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ccli-tz" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")

		viper.SetConfigName(config.FileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
