package leader

import (
	"os/exec"
)

type CmdRunner struct {
	DryRun bool
}

func (CmdRunner) GetSchedule(trimmedArgs []string) (string, error) {
	aCmd := exec.Command("cardano-cli", trimmedArgs...)
	stdout, err := aCmd.CombinedOutput()
	return string(stdout), err
}

func CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic string) []string {
	trimmedArgs := []string{"query", "leadership-schedule",
		"--vrf-signing-key-file", vrfKeysFile,
		"--stake-pool-id", poolID,
		"--genesis", shelleyGenesisFile,
		period,
	}
	if testnetMagic != "" {
		trimmedArgs = append(trimmedArgs, "--testnet-magic", testnetMagic)
	} else {
		trimmedArgs = append(trimmedArgs, "--mainnet")
	}
	return trimmedArgs
}
