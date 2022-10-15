package cmd

import (
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
	"github.com/lambda-honeypot/ccli-tz/pkg/server"
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the leaderlog in server mode",
	Long: `Runs the leaderlog in a server mode so that pre-calculated schedules can be accessed via http call. For example:

ccli-tz server --testnet-magic 1 --config ~/some/custom/config.yml

Will run a server that hosts the following endpoints:

curl localhost:8080/current # Displays the schedule for the current epoch
curl localhost:8080/next # Displays the schedule for the next epoch

WARNING: It is advised that you do not expose these endpoints externally to your block producing network. 
`,
	RunE: RunServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunServer(cmd *cobra.Command, _ []string) error {
	fmt.Println("server started")
	testnetMagic, err := cmd.Flags().GetString("testnet-magic")
	if err != nil {
		return fmt.Errorf("failed to get testnet string: %v", err)
	}
	cfg := config.ReadConfig()
	fileUtils := &utils.FileUtils{}
	err = server.WebServer(testnetMagic, &leader.CmdRunner{}, cfg, fileUtils)
	return err
}
