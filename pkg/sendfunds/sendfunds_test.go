package sendfunds

import (
	"github.com/golang/mock/gomock"
	mock_leader "github.com/lambda-honeypot/ccli-tz/pkg/leader/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSendFunds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test SendFunds")
}

var _ = Describe("createUTXOFromAddress", func() {
	var (
		ctrl       *gomock.Controller
		mockRunner *mock_leader.MockCommandRunner
		fundSender FundSender
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRunner = mock_leader.NewMockCommandRunner(ctrl)
		fundSender = NewFundSender(mockRunner, "net", "magic")
	})

	It("should produce the expected balances", func() {
		exampleTXs := `
TxHash                                 TxIx        Amount
--------------------------------------------------------------------------------------
392bd0b756aef669c2b17d706913d1ed263c5faea6005f088fbd65f0d1fa9d0f     0        1000069816722 lovelace + 982 2b59b3d42ed160682c95a8881e0956df7ae4979dc7cc5335b4702106.STPZcoin + TxOutDatumHashNon
`
		mockRunner.EXPECT().RunCardanoCmd(gomock.Any()).Return(exampleTXs, nil)
		utxo, err := fundSender.createUTXOFromAddress("something")
		Expect(err).ToNot(HaveOccurred())
		Expect(utxo.ADABalance).To(Equal(1000069816722))
		Expect(utxo.TXCount).To(Equal(1))
	})
})

var _ = Describe("generateTokenDetailsAndVerify", func() {

	It("should produce the expected balances with no tokens", func() {
		utxoDetails := generateTestUTXO(map[string]int{})
		paymentAddressesWithTokens := map[string]PaymentDetails{
			"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 21415881},
			"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 3717996},
			"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 33818574},
		}
		var expectedTokens []TokenDetails
		paymentTokenDetails, err := generateTokenDetailsAndVerify(utxoDetails, paymentAddressesWithTokens)
		Expect(err).ToNot(HaveOccurred())
		Expect(paymentTokenDetails).To(Equal(expectedTokens))
	})

	It("should produce the expected balances with one token", func() {
		token1 := "2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons"
		utxoDetails := generateTestUTXO(map[string]int{token1: 100})
		paymentAddressesWithTokens := map[string]PaymentDetails{
			"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 21415881, PaymentTokens: []TokenDetails{{token1, 9}}},
			"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 3717996, PaymentTokens: []TokenDetails{{token1, 3}}},
			"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 33818574, PaymentTokens: []TokenDetails{{token1, 6}}},
		}
		expectedTokens := []TokenDetails{{token1, 82}}
		paymentTokenDetails, err := generateTokenDetailsAndVerify(utxoDetails, paymentAddressesWithTokens)
		Expect(err).ToNot(HaveOccurred())
		Expect(paymentTokenDetails).To(Equal(expectedTokens))
	})

	It("should produce the expected balances with multiple tokens", func() {
		token1 := "2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons"
		token2 := "2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.MegaHoskinsons"
		token3 := "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY"
		utxoDetails := generateTestUTXO(map[string]int{token1: 100, token2: 50, token3: 80})
		paymentAddressesWithTokens := map[string]PaymentDetails{
			"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 21415881, PaymentTokens: []TokenDetails{
				{token1, 9},
				{token2, 13},
				{token3, 24},
			}},
			"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 3717996, PaymentTokens: []TokenDetails{
				{token1, 3},
				{token2, 15},
				{token3, 23},
			}},
			"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 33818574, PaymentTokens: []TokenDetails{
				{token1, 6},
				{token2, 16},
				{token3, 29},
			}},
		}
		expectedTokens := []TokenDetails{
			{token1, 82},
			{token2, 6},
			{token3, 4},
		}
		paymentTokenDetails, err := generateTokenDetailsAndVerify(utxoDetails, paymentAddressesWithTokens)
		Expect(err).ToNot(HaveOccurred())
		Expect(paymentTokenDetails).To(Equal(expectedTokens))
	})
})
