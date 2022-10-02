package leader

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CalculateLeaderArgs", func() {
	It("should only return non plank args", func() {
		period := "--current"
		shelleyGenesisFile := "shelley-genesis.json"
		poolID := "5be57ce6d1225697f4ad4090355f0a72d6e1e2446d1d768f36aa118c"
		vrfKeysFile := "vrf.skey"
		testnetMagic := ""
		resultArgs := CalculateLeaderArgs(period, shelleyGenesisFile, poolID, vrfKeysFile, testnetMagic)
		for _, arg := range resultArgs {
			Expect(arg).ToNot(Equal(""))
		}
	})
})
