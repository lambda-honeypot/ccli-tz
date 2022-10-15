package config

import (
	"fmt"
	"github.com/golang/mock/gomock"
	mocks "github.com/lambda-honeypot/ccli-tz/pkg/config/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

const sampleString = `VRFSigningKeyFile: /path/to/vrf-signing-key-file/vrf.skey
stakePoolID: <insert pool id>
shelleyGenesisFile: /path/to/genesis-file/shelley-genesis.json
timeZone: Europe/London
serverPort: "9091"
persistMode: true
`

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config Creating", func() {
	var (
		ctrl              *gomock.Controller
		mockConfigCreator *mocks.MockCfgCreator
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockConfigCreator = mocks.NewMockCfgCreator(ctrl)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Context("InitialiseConfigFile", func() {
		It("should attempt to create expected config file", func() {
			configFilePath := "/home/someone/.ccli-tz"
			sampleBytes := []byte(sampleString)
			mockConfigCreator.EXPECT().GetConfigFilePath().Return(configFilePath, nil)
			mockConfigCreator.EXPECT().ConfigFileExists(configFilePath).Return(false, nil)
			mockConfigCreator.EXPECT().WriteFile(configFilePath, sampleBytes)
			InitialiseConfigFile(mockConfigCreator)
		})
	})
	Context("writeTemplateConfig", func() {
		It("should attempt to create expected config file", func() {
			configFilePath := "/home/someone/.ccli-tz"
			sampleBytes := []byte(sampleString)
			mockConfigCreator.EXPECT().ConfigFileExists(configFilePath).Return(false, nil)
			mockConfigCreator.EXPECT().WriteFile(configFilePath, sampleBytes)
			err := writeTemplateConfig(configFilePath, mockConfigCreator)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should NOT attempt to create expected config file when exists", func() {
			configFilePath := "/home/someone/.ccli-tz"
			mockConfigCreator.EXPECT().ConfigFileExists(configFilePath).Return(true, nil)
			err := writeTemplateConfig(configFilePath, mockConfigCreator)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should fail with error when checking config file exists has error", func() {
			configFilePath := "/home/someone/.ccli-tz"
			mockConfigCreator.EXPECT().ConfigFileExists(configFilePath).Return(false, fmt.Errorf("failed reading: %s", configFilePath))
			err := writeTemplateConfig(configFilePath, mockConfigCreator)
			Expect(err).To(HaveOccurred())
		})
	})
})
