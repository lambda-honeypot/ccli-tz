package leader

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
	"time"
)

type CommandRunner interface {
	GetSchedule(period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic string, dryRun bool) (string, error)
}

type ScheduleRow struct {
	BlockCount    int
	SlotNumber    int
	ScheduledTime string
}

func CreateAndRun(args []string, testnetMagic string, cmdRunner CommandRunner, dryRun bool) error {
	period := "--" + args[0]
	shelleyGenesisFile := normaliseHomeDir(viper.GetString("shelleyGenesisFile"))
	vrfKeysFile := normaliseHomeDir(viper.GetString("VRFSigningKeyFile"))
	poolId := viper.GetString("stakePoolID")
	timeZone := viper.GetString("timeZone")
	fmt.Println(fmt.Sprintf("Calculating for pool: %s", poolId))
	schedule, err := CalcTZSchedule(timeZone, period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic, cmdRunner, dryRun)
	if dryRun {
		return nil
	}
	if err != nil {
		return err
	}
	PrintSchedule(schedule)
	return nil
}

func normaliseHomeDir(file string) string {
	if strings.HasPrefix(file, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Errorf("failed to get userhome dir with: %v", err)
			return file
		}
		return strings.Replace(file, "~", home, 1)
	}
	return file
}

func PrintSchedule(schedule []ScheduleRow) {
	b, err := json.MarshalIndent(schedule, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func CalcTZSchedule(timeZone, period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic string, runner CommandRunner, dryRun bool) ([]ScheduleRow, error) {
	var rows []ScheduleRow
	rawSchedule, err := runner.GetSchedule(period, shelleyGenesisFile, poolId, vrfKeysFile, testnetMagic, dryRun)
	if dryRun {
		return rows, nil
	}
	if err != nil {
		log.Errorf("failed to run leader log with: %s", rawSchedule)
		return rows, err
	}
	lines := strings.Split(rawSchedule, "\n")
	for i, line := range lines[2:] {
		rawSpaceSplit := strings.Split(strings.TrimSpace(line), "  ")
		var spaceSplit = []string{}
		for _, elem := range rawSpaceSplit {
			if strings.TrimSpace(elem) != "" {
				spaceSplit = append(spaceSplit, elem)
			}
		}
		if len(spaceSplit) != 2 {
			continue
		}
		rawTS := strings.TrimSpace(spaceSplit[len(spaceSplit)-1])
		convertedTime, err := convertTime(rawTS, timeZone)
		if err != nil {
			log.Errorf("failed to convert time with: %v", err)
		}
		slot, err := strconv.Atoi(strings.TrimSpace(spaceSplit[0]))
		if err != nil {
			log.Errorf("failed to convert slot num with: %v", err)
		}
		row := ScheduleRow{BlockCount: i + 1, SlotNumber: slot, ScheduledTime: convertedTime}
		rows = append(rows, row)
	}
	return rows, nil
}

func convertTime(timeStamp, timeZone string) (string, error) {
	layout := "2006-01-02 15:04:05 MST"
	t, err := time.Parse(layout, timeStamp)
	if err != nil {
		return timeStamp, err
	}
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return timeStamp, err
	}
	return t.In(loc).Format(layout), nil
}
