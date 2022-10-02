package leader

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TipData struct {
	Block        int    `json:"block"`
	Epoch        int    `json:"epoch"`
	Era          string `json:"era"`
	Hash         string `json:"hash"`
	Slot         int    `json:"slot"`
	SyncProgress string `json:"syncProgress"`
}

func GetTipData(testnetMagic string, runner CommandRunner) (*TipData, error) {
	tipArgs := CalculateTipArgs(testnetMagic)
	rawTip, err := runner.RunCardanoCmd(tipArgs)
	if err != nil {
		log.Errorf("failed to get tip with: %s", rawTip)
		return nil, err
	}
	tipData := &TipData{}
	err = json.Unmarshal([]byte(rawTip), tipData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall tip data with: %v", err)
	}
	return tipData, nil
}
