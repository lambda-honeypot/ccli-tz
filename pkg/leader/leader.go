package leader

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type CommandRunner interface {
	GetSchedule(period, shelleyGenesisFile, poolId, networkMagic, vrfKeysFile string) (string, error)
}

type ScheduleRow struct {
	SlotNumber int
	TimeZone string
}

func CreateAndRun(timeZone string, cmdRunner CommandRunner) {
	schedule := CalcTZSchedule(timeZone, cmdRunner)
	PrintSchedule(schedule)
}

func PrintSchedule(schedule []ScheduleRow) {
	b, err := json.MarshalIndent(schedule, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func CalcTZSchedule(timeZone string, runner CommandRunner) []ScheduleRow {
	var rows []ScheduleRow
	rawSchedule, _ := runner.GetSchedule()
	lines := strings.Split(rawSchedule, "\n")
	for _, line := range lines[2:] {
		spaceSplit := strings.Split(strings.TrimSpace(line), "  ")
		rawTS := strings.TrimSpace(spaceSplit[len(spaceSplit)-1])
		convertedTime, err := convertTime(rawTS, timeZone)
		if err != nil {
			log.Errorf("failed with err: %v", err)
		}
		//newLine := strings.TrimSpace(spaceSplit[0]) + " " + convertedTime + "\n"
		slot, err := strconv.Atoi(strings.TrimSpace(spaceSplit[0]))
		if err != nil {
			log.Errorf("failed with err: %v", err)
		}
		row := ScheduleRow{SlotNumber: slot, TimeZone: convertedTime}
		rows = append(rows, row)
	}
	return rows
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
