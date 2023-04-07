package sendfunds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FullUTXO", func() {
	It("should parse output correctly with tokens", func() {
		exampleUTXO := `		
						   TxHash                                 TxIx        Amount
--------------------------------------------------------------------------------------
ae7828ff0c6d607a279aba9e483617fc5ea423b1bf890bbe1b093007c791cadf     0        4015245989 lovelace + 9062597 2aa9c1557fcf8e7caa049fa0911a8724a1cdaf8037fe0b431c6ac664.PIGYToken + TxOutDatumHashNone
`
		utxo, err := parseFullUTXO(exampleUTXO, "addr1_somehash")
		Expect(err).To(Not(HaveOccurred()))
		Expect(len(utxo.Rows)).To(Equal(1))
		Expect(utxo.Address).To(Equal("addr1_somehash"))
		Expect(utxo.TXCount).To(Equal(1))
		Expect(utxo.ADABalance).To(Equal(4015245989))

		firstRow := utxo.Rows[0]
		Expect(firstRow.Hash).To(Equal("ae7828ff0c6d607a279aba9e483617fc5ea423b1bf890bbe1b093007c791cadf"))
		Expect(firstRow.TxID).To(Equal("0"))
		Expect(firstRow.LovelaceBalance).To(Equal(4015245989))

		Expect(len(firstRow.Tokens)).To(Equal(1))
		Expect(firstRow.Tokens["2aa9c1557fcf8e7caa049fa0911a8724a1cdaf8037fe0b431c6ac664.PIGYToken"]).To(Equal(9062597))
	})

	It("should parse output correctly without tokens", func() {
		exampleUTXO := `
                           TxHash                                 TxIx        Amount
--------------------------------------------------------------------------------------
b5dfee5e9562cdaff11b9de28ee404d4ceb6d4617f5e90180966df2b12d77ac8     0        1520359733 lovelace + TxOutDatumHashNone
`
		utxo, err := parseFullUTXO(exampleUTXO, "addr2_somehash")
		Expect(err).To(Not(HaveOccurred()))
		Expect(utxo.Address).To(Equal("addr2_somehash"))
		Expect(len(utxo.Rows)).To(Equal(1))
		Expect(utxo.TXCount).To(Equal(1))
		Expect(utxo.ADABalance).To(Equal(1520359733))

		firstRow := utxo.Rows[0]
		Expect(firstRow.Hash).To(Equal("b5dfee5e9562cdaff11b9de28ee404d4ceb6d4617f5e90180966df2b12d77ac8"))
		Expect(firstRow.TxID).To(Equal("0"))
		Expect(firstRow.LovelaceBalance).To(Equal(1520359733))
	})
	//revive:disable:line-length-limit
	It("should parse output correctly with many tokens", func() {
		utxoString := `                           TxHash                                 TxIx        Amount
--------------------------------------------------------------------------------------
12a04dd49c694d99d612c0128033743c470587e59a40584c7ce7ba921f150c42     1        6407284 lovelace + 1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons + 10 3e07d877c7115f8b402de4a1c00ad9c1f0f00a9a1c25bec36c4cb634.LOVE + 1 44ae79bd8fbbcd4b7ecbf5745feaa75feb4c7be8ac5198a968528f0a.HoskPoolOG + 1 4fc17ed958dbf8ff4f2f7e876b262ceb8c127511ec99419c65f3cd6f.Cryptodelics005 + 1 639936cfb38a4dd74bfd261795e98f6eb9bac8d9e0c72cff6d7245fe.SpacePopsJupiter48 + 1 6ca288470a332a534a789f607aa841e1176473a380e131dbbb45012b.MosasaurusRegular01 + 845575 98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.hosk + 179 98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.kilohosk + 859905 98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.microhosk + 498993 98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.millihosk + 1 a0391188e8a83bcbc51a2ed575551532db898e4d94b2a8f43825c74b.PiratesHoskinsonCCA + 25000 bb034f986a145377d4767a2dbbadaaa630bb76c741c8efc6479dbb9e.Schmeckle + 2 bfad036b5bc09ac5854faa58c9a6c6d1c0c87468052cfcd4bd832e24.testcoin4 + 1 cd27bd963a1bab843b55b787286a25739f6d687a810c3a92df79c735.DYOR037 + 2140 d45b8c3932daa3e7cc0148553e89cb93f83fbeb34d6b4b8b712ccc9f.withspacesCoin + TxOutDatumNone `
		utxo, err := parseFullUTXO(utxoString, "addr5_somehash")
		Expect(err).To(Not(HaveOccurred()))
		Expect(utxo.Address).To(Equal("addr5_somehash"))
		Expect(len(utxo.Rows)).To(Equal(1))
		Expect(utxo.TXCount).To(Equal(1))
		Expect(utxo.ADABalance).To(Equal(6407284))

		Expect(len(utxo.TokenBalances)).To(Equal(15))
		Expect(utxo.TokenBalances["2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons"]).To(Equal(1))
		Expect(utxo.TokenBalances["3e07d877c7115f8b402de4a1c00ad9c1f0f00a9a1c25bec36c4cb634.LOVE"]).To(Equal(10))
		Expect(utxo.TokenBalances["44ae79bd8fbbcd4b7ecbf5745feaa75feb4c7be8ac5198a968528f0a.HoskPoolOG"]).To(Equal(1))
		Expect(utxo.TokenBalances["4fc17ed958dbf8ff4f2f7e876b262ceb8c127511ec99419c65f3cd6f.Cryptodelics005"]).To(Equal(1))
		Expect(utxo.TokenBalances["639936cfb38a4dd74bfd261795e98f6eb9bac8d9e0c72cff6d7245fe.SpacePopsJupiter48"]).To(Equal(1))
		Expect(utxo.TokenBalances["6ca288470a332a534a789f607aa841e1176473a380e131dbbb45012b.MosasaurusRegular01"]).To(Equal(1))
		Expect(utxo.TokenBalances["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.hosk"]).To(Equal(845575))
		Expect(utxo.TokenBalances["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.kilohosk"]).To(Equal(179))
		Expect(utxo.TokenBalances["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.microhosk"]).To(Equal(859905))
		Expect(utxo.TokenBalances["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.millihosk"]).To(Equal(498993))
		Expect(utxo.TokenBalances["a0391188e8a83bcbc51a2ed575551532db898e4d94b2a8f43825c74b.PiratesHoskinsonCCA"]).To(Equal(1))
		Expect(utxo.TokenBalances["bb034f986a145377d4767a2dbbadaaa630bb76c741c8efc6479dbb9e.Schmeckle"]).To(Equal(25000))
		Expect(utxo.TokenBalances["bfad036b5bc09ac5854faa58c9a6c6d1c0c87468052cfcd4bd832e24.testcoin4"]).To(Equal(2))
		Expect(utxo.TokenBalances["cd27bd963a1bab843b55b787286a25739f6d687a810c3a92df79c735.DYOR037"]).To(Equal(1))
		Expect(utxo.TokenBalances["d45b8c3932daa3e7cc0148553e89cb93f83fbeb34d6b4b8b712ccc9f.withspacesCoin"]).To(Equal(2140))

		firstRow := utxo.Rows[0]
		Expect(firstRow.Hash).To(Equal("12a04dd49c694d99d612c0128033743c470587e59a40584c7ce7ba921f150c42"))
		Expect(firstRow.TxID).To(Equal("1"))
		Expect(firstRow.LovelaceBalance).To(Equal(6407284))

		Expect(len(firstRow.Tokens)).To(Equal(15))
		Expect(firstRow.Tokens["2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons"]).To(Equal(1))
		Expect(firstRow.Tokens["3e07d877c7115f8b402de4a1c00ad9c1f0f00a9a1c25bec36c4cb634.LOVE"]).To(Equal(10))
		Expect(firstRow.Tokens["44ae79bd8fbbcd4b7ecbf5745feaa75feb4c7be8ac5198a968528f0a.HoskPoolOG"]).To(Equal(1))
		Expect(firstRow.Tokens["4fc17ed958dbf8ff4f2f7e876b262ceb8c127511ec99419c65f3cd6f.Cryptodelics005"]).To(Equal(1))
		Expect(firstRow.Tokens["639936cfb38a4dd74bfd261795e98f6eb9bac8d9e0c72cff6d7245fe.SpacePopsJupiter48"]).To(Equal(1))
		Expect(firstRow.Tokens["6ca288470a332a534a789f607aa841e1176473a380e131dbbb45012b.MosasaurusRegular01"]).To(Equal(1))
		Expect(firstRow.Tokens["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.hosk"]).To(Equal(845575))
		Expect(firstRow.Tokens["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.kilohosk"]).To(Equal(179))
		Expect(firstRow.Tokens["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.microhosk"]).To(Equal(859905))
		Expect(firstRow.Tokens["98dc68b04026544619a251bc01aad2075d28433524ac36cbc75599a1.millihosk"]).To(Equal(498993))
		Expect(firstRow.Tokens["a0391188e8a83bcbc51a2ed575551532db898e4d94b2a8f43825c74b.PiratesHoskinsonCCA"]).To(Equal(1))
		Expect(firstRow.Tokens["bb034f986a145377d4767a2dbbadaaa630bb76c741c8efc6479dbb9e.Schmeckle"]).To(Equal(25000))
		Expect(firstRow.Tokens["bfad036b5bc09ac5854faa58c9a6c6d1c0c87468052cfcd4bd832e24.testcoin4"]).To(Equal(2))
		Expect(firstRow.Tokens["cd27bd963a1bab843b55b787286a25739f6d687a810c3a92df79c735.DYOR037"]).To(Equal(1))
		Expect(firstRow.Tokens["d45b8c3932daa3e7cc0148553e89cb93f83fbeb34d6b4b8b712ccc9f.withspacesCoin"]).To(Equal(2140))
	})
	//revive:enable:line-length-limit
})
