package leader

import (
	"os/exec"
)

type CmdRunner struct{}

func (CmdRunner) RunCardanoCmd(trimmedArgs []string) (string, error) {
	aCmd := exec.Command("cardano-cli", trimmedArgs...)
	stdout, err := aCmd.CombinedOutput()
	return string(stdout), err
}

func CalculateLeaderArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic string) []string {
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

func CalculateTipArgs(testnetMagic string) []string {
	trimmedArgs := []string{"query", "tip"}
	if testnetMagic != "" {
		trimmedArgs = append(trimmedArgs, "--testnet-magic", testnetMagic)
	} else {
		trimmedArgs = append(trimmedArgs, "--mainnet")
	}
	return trimmedArgs
}
