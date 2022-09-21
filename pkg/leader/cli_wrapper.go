package leader

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type CmdRunner struct {
}

func (r CmdRunner) GetSchedule(period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic string, dryRun bool) (string, error) {
	trimmedArgs := []string{"query", "leadership-schedule",
		"--vrf-signing-key-file", vrfKeysFile,
		"--stake-pool-id", poolId,
		"--genesis", shelleyGenesisFile,
		period,
	}
	if testnetMagic != "" {
		trimmedArgs = append(trimmedArgs, "--testnet-magic", testnetMagic)
	} else {
		trimmedArgs = append(trimmedArgs, "--mainnet", testnetMagic)
	}
	if dryRun {
		log.Infof("dry-run, would have executed:\n\ncardano-cli %v", trimmedArgs)
		return "", nil
	}
	aCmd := exec.Command("cardano-cli", trimmedArgs...)
	stdout, err := aCmd.CombinedOutput()
	return string(stdout), err
}
