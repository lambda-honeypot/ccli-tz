package leader

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/lambda-honeypot/ccli-tz/pkg/config"
	mocks "github.com/lambda-honeypot/ccli-tz/pkg/leader/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

const rawSlotOutput = `     SlotNo                          UTC Time              
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
     71425763                   2022-09-12 14:14:14 UTC
`

const emptySlotOutput = `SlotNo                          UTC Time
-------------------------------------------------------------`

func TestLeader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Leader Suite")
}

var _ = Describe("CreateAndRun", func() {
	var (
		ctrl              *gomock.Controller
		mockCommandRunner *mocks.MockCommandRunner
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockCommandRunner = mocks.NewMockCommandRunner(ctrl)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Context("CreateAndRun", func() {
		It("should not error when called with valid arguments", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			vrfKeysFile := "vrf.skey"
			timeZone := "America/New_York"
			cfg := &config.CfgYaml{
				VRFSigningKeyFile: vrfKeysFile,
				StakePoolID:       poolID,
				GenesisFile:       shelleyGenesisFile,
				TimeZone:          timeZone,
			}
			testnetMagic := ""
			trimmedArgs := CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
			mockCommandRunner.EXPECT().GetSchedule(trimmedArgs).Return(rawSlotOutput, nil)
			args := []string{"current"}

			err := CreateAndRun(args, testnetMagic, mockCommandRunner, cfg)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should not error when called with valid arguments and dry run", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			vrfKeysFile := "vrf.skey"
			timeZone := "America/New_York"
			cfg := &config.CfgYaml{
				VRFSigningKeyFile: vrfKeysFile,
				StakePoolID:       poolID,
				GenesisFile:       shelleyGenesisFile,
				TimeZone:          timeZone,
			}
			testnetMagic := ""
			trimmedArgs := CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
			mockCommandRunner.EXPECT().GetSchedule(trimmedArgs).Return(rawSlotOutput, nil)
			args := []string{"current"}

			err := CreateAndRun(args, testnetMagic, mockCommandRunner, cfg)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should error when GetSchedule returns error", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			vrfKeysFile := "vrf.skey"
			timeZone := "America/New_York"
			cfg := &config.CfgYaml{
				VRFSigningKeyFile: vrfKeysFile,
				StakePoolID:       poolID,
				GenesisFile:       shelleyGenesisFile,
				TimeZone:          timeZone,
			}
			testnetMagic := ""
			trimmedArgs := CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
			mockCommandRunner.EXPECT().GetSchedule(trimmedArgs).Return("", errors.New("waah"))
			args := []string{"current"}

			err := CreateAndRun(args, testnetMagic, mockCommandRunner, cfg)
			Expect(err).To(HaveOccurred())
		})
	})
	Context("CalcTZSchedule", func() {
		It("should calculate the correct rows with sample cardano rows", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			vrfKeysFile := "vrf.skey"
			timeZone := "America/Guatemala"
			testnetMagic := ""
			trimmedArgs := CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
			mockCommandRunner.EXPECT().GetSchedule(trimmedArgs).Return(rawSlotOutput, nil)

			schedule, err := CalcTZSchedule(timeZone, trimmedArgs, mockCommandRunner)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(schedule)).To(Equal(11))
			row1 := ScheduleRow{
				BlockCount:    1,
				SlotNumber:    71029049,
				ScheduledTime: "2022-09-07 18:02:20 CST",
			}
			Expect(schedule[0]).To(Equal(row1))
			row5 := ScheduleRow{
				BlockCount:    5,
				SlotNumber:    71226203,
				ScheduledTime: "2022-09-10 00:48:14 CST",
			}
			Expect(schedule[4]).To(Equal(row5))
			row11 := ScheduleRow{
				BlockCount:    11,
				SlotNumber:    71425763,
				ScheduledTime: "2022-09-12 08:14:14 CST",
			}
			Expect(schedule[10]).To(Equal(row11))
		})
		It("should calculate the correct rows with sample cardano empty rows output", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			vrfKeysFile := "vrf.skey"
			timeZone := "America/Guatemala"
			testnetMagic := ""
			trimmedArgs := CalculateArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
			mockCommandRunner.EXPECT().GetSchedule(trimmedArgs).Return(emptySlotOutput, nil)

			schedule, err := CalcTZSchedule(timeZone, trimmedArgs, mockCommandRunner)
			Expect(err).ToNot(HaveOccurred())
			Expect(len(schedule)).To(Equal(0))
		})
	})
	Context("ScheduleOutputString", func() {
		It("should generate correct output string when there are rows", func() {
			row1 := ScheduleRow{
				BlockCount:    1,
				SlotNumber:    71029049,
				ScheduledTime: "2022-09-07 18:02:20 CST",
			}
			row2 := ScheduleRow{
				BlockCount:    2,
				SlotNumber:    71226203,
				ScheduledTime: "2022-09-10 00:48:14 CST",
			}
			row3 := ScheduleRow{
				BlockCount:    3,
				SlotNumber:    71425763,
				ScheduledTime: "2022-09-12 08:14:14 CST",
			}
			period := "next"
			schedule := []ScheduleRow{row1, row2, row3}
			outputString := ScheduleOutputString(schedule, period)
			expectedOutput := `[
  {
    "BlockCount": 1,
    "SlotNumber": 71029049,
    "ScheduledTime": "2022-09-07 18:02:20 CST"
  },
  {
    "BlockCount": 2,
    "SlotNumber": 71226203,
    "ScheduledTime": "2022-09-10 00:48:14 CST"
  },
  {
    "BlockCount": 3,
    "SlotNumber": 71425763,
    "ScheduledTime": "2022-09-12 08:14:14 CST"
  }
]`
			Expect(outputString).To(Equal(expectedOutput))
		})
		It("should generate correct output string when there are no rows for current", func() {
			period := "current"
			var schedule []ScheduleRow
			outputString := ScheduleOutputString(schedule, period)
			Expect(outputString).To(Equal("No schedule blocks for current epoch"))
		})
		It("should generate correct output string when there are no rows for next", func() {
			period := "next"
			var schedule []ScheduleRow
			outputString := ScheduleOutputString(schedule, period)
			Expect(outputString).To(Equal("No schedule blocks for next epoch"))
		})
	})
})
