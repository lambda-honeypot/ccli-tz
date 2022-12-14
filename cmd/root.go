package cmd

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
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

ccli-tz next

This will create the leadership-schedule for the pool and timezone referenced within the config file for the next epoch.`,
	Args: cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE:    LeadershipLog,
	PreRunE: initConfig,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func LeadershipLog(command *cobra.Command, args []string) error {
	err := validateArgs(args)
	if err != nil {
		return err
	}
	testnetMagic, err := command.Flags().GetString("testnet-magic")
	if err != nil {
		return fmt.Errorf("failed to get testnet string: %v", err)
	}
	dryRun, err := command.Flags().GetBool("dry-run")
	if err != nil {
		return err
	}
	cfg := config.ReadConfig()
	if dryRun {
		leader.LogOutParams(args, testnetMagic, cfg)
	} else {
		err = leader.CreateAndRun(args, testnetMagic, &leader.CmdRunner{}, cfg)
	}
	if err != nil {
		return fmt.Errorf("failed to run leadership log with: %v", err)
	}
	return nil
}

func validateArgs(args []string) error {
	validValues := []string{"current", "next"}
	if len(args) != 1 {
		return fmt.Errorf("incorrect number of args. Got %d expected 1", len(args))
	}
	for _, value := range validValues {
		if args[0] == value {
			return nil
		}
	}
	return fmt.Errorf("incorrect arg supplied. Got `%s` but expect one of %v", args[0], validValues)
}

func init() {
	//cobra.OnInitialize(initConfig)

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
func initConfig(_ *cobra.Command, _ []string) error {
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
	var err error
	if err = viper.ReadInConfig(); err == nil {
		_, err = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		return err
	}
	return fmt.Errorf("failed to read config file with err: %v", err)
}
