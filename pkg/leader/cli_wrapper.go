package leader

import "os/exec"

type CmdRunner struct {
}

func (r CmdRunner) GetSchedule(period, shelleyGenesisFile, poolId, networkMagic, vrfKeysFile string) (string, error) {
	//period := "--current"
	//shelleyGenesisFile := "/var/lib/cardano/mainnet-shelley-genesis.json"
	//poolId := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
	//networkMagic := "--mainnet"
	//vrfKeysFile := "vrf.skey"
	trimmedArgs := []string{"query", "leadership-schedule",
		"--vrf-signing-key-file", vrfKeysFile,
		"--stake-pool-id", poolId,
		"--genesis", shelleyGenesisFile,
		period,
		networkMagic,
	}

	aCmd := exec.Command("cardano-cli", trimmedArgs...)
	stdout, err := aCmd.CombinedOutput()
	return string(stdout), err
}

