package cmd

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
	"github.com/lambda-honeypot/ccli-tz/pkg/sendfunds"
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// serverCmd represents the server command
var sendFundCmd = &cobra.Command{
	Use:   "sendfunds",
	Short: "Run the sendfunds command",
	Long: `Runs the sendfunds in a server mode so that pre-calculated schedules can be accessed via http call. For example:

# SIGNING_KEY_FILE must be specified as an environment variable
export SIGNING_KEY_FILE=/path/to/source/payment.skey 
ccli-tz sendfunds --payment-file ~/some/path/to/payment.yml --testnet-magic 1

# mainnet requires no flag
SIGNING_KEY_FILE=/path/to/source/payment.skey ccli-tz sendfunds --payment-file ~/some/path/to/payment.yml 

This will 
`,
	RunE: RunSendFunds,
}

func init() {
	rootCmd.AddCommand(sendFundCmd)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	sendFundCmd.Flags().String("payment-file", "", "If set to true will print the command and args passed to cardano-cli")
}

func RunSendFunds(cmd *cobra.Command, _ []string) error {
	log.SetLevel(log.DebugLevel)
	runner := leader.CmdRunner{}
	testnetMagic, err := cmd.Flags().GetString("testnet-magic")
	network := "--mainnet"
	if testnetMagic != "" {
		network = "--testnet-magic"
	}
	paymentFilePath, err := cmd.Flags().GetString("payment-file")
	if err != nil {
		return fmt.Errorf("failed to get --payment-file path with: %v", err)
	}
	signingKeyFile := os.Getenv("SIGNING_KEY_FILE")
	if signingKeyFile == "" {
		return fmt.Errorf("error SIGNING_KEY_FILE variable is not set")
	}

	fileUtils := &utils.FileUtils{}
	paymentYaml := sendfunds.ReadPaymentFile(fileUtils, paymentFilePath)
	paymentAddressesWithTokens := paymentYaml.TargetAddresses
	startAddress := paymentYaml.SourceAddress

	fs := sendfunds.NewFundSender(runner, network, testnetMagic)
	return fs.RunSendFunds(startAddress, signingKeyFile, paymentAddressesWithTokens)
}
