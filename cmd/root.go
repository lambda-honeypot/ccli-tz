package cmd

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var testnet string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccli-tz [next | current]",
	Short: "Wrapper around the cardano-cli leadership schedule with timezone conversion",
	Long: `Wraps around the cardano-cli query leadership-schedule and allows for timezone conversion of dates. For example:

ccli-tz --timezone europe/london next

This will create the leadership-schedule for the pool and keys within the config file for the next epoch.`,
	Args: cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: LeadershipLog,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func LeadershipLog(command *cobra.Command, args []string) {
	validateArgs(args)
	testnetMagic, err := command.Flags().GetString("testnet-magic")
	if err != nil {
		log.Fatalf("failed to get testnet string: %v", err)
	}
	dryRun, err := command.Flags().GetBool("dry-run")

	err = leader.CreateAndRun(args, testnetMagic, &leader.CmdRunner{}, dryRun)
	if err != nil {
		log.Fatalf("failed to run leadership log with: %v", err)
	}
}

func validateArgs(args []string) {
	validValues := []string{"current", "next"}
	if len(args) != 1 {
		log.Fatalf("incorrect number of args. Got %d expected 1", len(args))
	}
	for _, value := range validValues {
		if args[0] == value {
			return
		}
	}
	log.Fatalf("incorrect arg supplied. Got `%s` but expect one of %v", args[0], validValues)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccli-tz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(&testnet, "testnet-magic", "", "Specify a testnet instead of mainnet")
	rootCmd.Flags().Bool("dry-run", false, "If set to true will print the command and args passed to cardano-cli")
	//rootCmd.Flags().Bool("current", false, "Calculate leader log for the current epoch")
	//rootCmd.Flags().BoolP("timezone", "t", false, "Help message for toggle")
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
	} else {
		log.Fatalf("failed to read config file at %s", config.FileName)
	}
}
