package leader

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

type CommandRunner interface {
	GetSchedule(command string) (string, error)
}

type ScheduleConvertor struct {
	CommandRunner CommandRunner
}

func LeadershipLog(_ *cobra.Command, _ []string) {
	CreateAndRun()
}

func CreateAndRun() {
	cmdRunner := &CmdRun{}
	sc := &ScheduleConvertor{cmdRunner}
	timeZone := "America/New_York"
	schedule := sc.CalcTZSchedule(timeZone)
	fmt.Println(schedule)
}

func (sc *ScheduleConvertor) CalcTZSchedule(timeZone string) string {
	var returnStr string
	rawSchedule, _ := sc.CommandRunner.GetSchedule("")
	lines := strings.Split(rawSchedule, "\n")
	for _, line := range lines[2:] {
		spaceSplit := strings.Split(strings.TrimSpace(line), "  ")
		rawTS := strings.TrimSpace(spaceSplit[len(spaceSplit)-1])
		convertedTime, err := convertTime(rawTS, timeZone)
		if err != nil {
			log.Errorf("failed with err: %v", err)
		}
		newLine := strings.TrimSpace(spaceSplit[0]) + " " + convertedTime
		returnStr += newLine
	}
	return returnStr
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

type CmdRun struct {
}

func (*CmdRun) GetSchedule(_ string) (string, error) {
	return `     SlotNo                          UTC Time              
-------------------------------------------------------------
     71029049                   2022-09-08 00:02:20 UTC
     71102016                   2022-09-08 20:18:27 UTC
     71108282                   2022-09-08 22:02:53 UTC
     71223290                   2022-09-10 05:59:41 UTC
     71226203                   2022-09-10 06:48:14 UTC
     71267198                   2022-09-10 18:11:29 UTC
     71351113                   2022-09-11 17:30:04 UTC
     71416799                   2022-09-12 11:44:50 UTC
     71419149                   2022-09-12 12:24:00 UTC
     71422743                   2022-09-12 13:23:54 UTC
     71425763                   2022-09-12 14:14:14 UTC`, nil
}
