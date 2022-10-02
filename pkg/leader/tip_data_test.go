package leader

import (
	"github.com/golang/mock/gomock"
	mocks "github.com/lambda-honeypot/ccli-tz/pkg/leader/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const sampleTip = `{
    "block": 7812478,
    "epoch": 365,
    "era": "Babbage",
    "hash": "6619890b976c9590bfe67a8fa2d59e857c32778eeec049d3adc4daefc5feb7c8",
    "slot": 72743242,
    "syncProgress": "100.00"
}
`

const unsyncTip = `{
    "block": 7812478,
    "epoch": 362,
    "era": "Babbage",
    "hash": "6619890b976c9590bfe67a8fa2d59e857c32778eeec049d3adc4daefc5feb7c8",
    "slot": 72743242,
    "syncProgress": "78.00"
}
`

var _ = Describe("GetTipData", func() {
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
	It("should parse sample correctly when called with valid args mainnet", func() {
		tipArgs := []string{"query", "tip", "--mainnet"}
		mockCommandRunner.EXPECT().RunCardanoCmd(tipArgs).Return(sampleTip, nil)
		tipData, err := GetTipData("", mockCommandRunner)
		Expect(err).ToNot(HaveOccurred())
		Expect(tipData.Epoch).To(Equal(365))
		Expect(tipData.Block).To(Equal(7812478))
		Expect(tipData.Era).To(Equal("Babbage"))
		Expect(tipData.Hash).To(Equal("6619890b976c9590bfe67a8fa2d59e857c32778eeec049d3adc4daefc5feb7c8"))
		Expect(tipData.Slot).To(Equal(72743242))
		Expect(tipData.SyncProgress).To(Equal("100.00"))
	})
	It("should parse sample correctly when called with valid args testnet", func() {
		tipArgs := []string{"query", "tip", "--testnet-magic", "1"}
		mockCommandRunner.EXPECT().RunCardanoCmd(tipArgs).Return(sampleTip, nil)
		tipData, err := GetTipData("1", mockCommandRunner)
		Expect(err).ToNot(HaveOccurred())
		Expect(tipData.Epoch).To(Equal(365))
		Expect(tipData.Block).To(Equal(7812478))
		Expect(tipData.Era).To(Equal("Babbage"))
		Expect(tipData.Hash).To(Equal("6619890b976c9590bfe67a8fa2d59e857c32778eeec049d3adc4daefc5feb7c8"))
		Expect(tipData.Slot).To(Equal(72743242))
		Expect(tipData.SyncProgress).To(Equal("100.00"))
	})
})
