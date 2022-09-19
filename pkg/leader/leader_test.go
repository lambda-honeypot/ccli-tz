package leader

import (
	"github.com/golang/mock/gomock"
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
     71425763                   2022-09-12 14:14:14 UTC`

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
	Context("InitialiseConfigFile", func() {
		It("should attempt to create expected config file", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolId := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			networkMagic := "--mainnet"
			vrfKeysFile := "vrf.skey"
			mockCommandRunner.EXPECT().GetSchedule(period, shelleyGenesisFile, poolId, networkMagic, vrfKeysFile).Return(rawSlotOutput, nil)
			timeZone := "America/New_York"
			CreateAndRun(timeZone, mockCommandRunner)
		})
	})
	Context("CalcTZSchedule", func() {
		It("should attempt to create expected config file", func() {
			period := "--current"
			shelleyGenesisFile := "shelley-genesis.json"
			poolId := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
			networkMagic := "--mainnet"
			vrfKeysFile := "vrf.skey"
			mockCommandRunner.EXPECT().GetSchedule(period, shelleyGenesisFile, poolId, networkMagic, vrfKeysFile).Return(rawSlotOutput, nil)
			timeZone := "America/New_York"
			schedule := CalcTZSchedule(timeZone, mockCommandRunner)
			Expect(schedule).To(Equal(""))
		})
	})
})
