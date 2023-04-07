package sendfunds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Build transaction Payment No Tokens", func() {
	expectedResponseManyPaymentsNoSendTokens := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
		"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911",
		"--tx-out", "addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk+3108141",
		"--tx-out", "addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl+1000000",
		"--tx-out", "addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd+5460362",
		"--tx-out", "addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su+74355476",
		"--tx-out", "addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l+1146843",
		"--tx-out", "addr1qx3g0pnxmmggj7jxckf064wzlfqjnyn9n4a2g7hn23wnjwcyzrx8ppcedh89pmnxdjm2hl4730eu38vzmfahxf22tgxs5hq8ja+7555538",
		"--tx-out", "addr1qxdkv4x590q00jtkd2pvhg5uejfgc29pxafj6wvelkl5msrk3hgcfqcwtf8j8yss6h6elqdeaerr8w89062u7v220tuqyqun8v+1000000",
		"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
		"--tx-out", "addr1qxlqxcxxn88fnf37wtgq8jwq8mfzumhq7svsnnr9crh0gkp0pxyg33eksp7m0p4atqcye7e3fzlgemnmwhcp26lwk4lq0fl74s+1422183",
		"--tx-out", "addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl+3584514",
		"--tx-out", "addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z+1000000",
		"--tx-out", "addr1qysfnqjx7rpg2l8g8ynjefwly5uftu5anqj3ze6939ywewm9nlxjsyf63x8tkxjagd70lea98mawg80k8lvjv2zma0fsq7xzls+1000000",
		"--invalid-hereafter", "48250738", "--fee", "239685", "--out-file", "/tmp/pay_multi3100205435/tx.raw"}

	expectedResponseManyPaymentsNoSendTokensButPaymentAddrHasSourceTokens := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
		"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+1 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY+1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
		"--tx-out", "addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk+3108141",
		"--tx-out", "addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl+1000000",
		"--tx-out", "addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd+5460362",
		"--tx-out", "addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su+74355476",
		"--tx-out", "addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l+1146843",
		"--tx-out", "addr1qx3g0pnxmmggj7jxckf064wzlfqjnyn9n4a2g7hn23wnjwcyzrx8ppcedh89pmnxdjm2hl4730eu38vzmfahxf22tgxs5hq8ja+7555538",
		"--tx-out", "addr1qxdkv4x590q00jtkd2pvhg5uejfgc29pxafj6wvelkl5msrk3hgcfqcwtf8j8yss6h6elqdeaerr8w89062u7v220tuqyqun8v+1000000",
		"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
		"--tx-out", "addr1qxlqxcxxn88fnf37wtgq8jwq8mfzumhq7svsnnr9crh0gkp0pxyg33eksp7m0p4atqcye7e3fzlgemnmwhcp26lwk4lq0fl74s+1422183",
		"--tx-out", "addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl+3584514",
		"--tx-out", "addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z+1000000",
		"--tx-out", "addr1qysfnqjx7rpg2l8g8ynjefwly5uftu5anqj3ze6939ywewm9nlxjsyf63x8tkxjagd70lea98mawg80k8lvjv2zma0fsq7xzls+1000000",
		"--invalid-hereafter", "48250738", "--fee", "239685", "--out-file", "/tmp/pay_multi3100205435/tx.raw"}

	manyPaymentAddresses := map[string]PaymentDetails{
		"addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah": {AdaAmount: 1253293},
		"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 1000000},
		"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 3584514},
		"addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su": {AdaAmount: 74355476},
		"addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk": {AdaAmount: 3108141},
		"addr1qxdkv4x590q00jtkd2pvhg5uejfgc29pxafj6wvelkl5msrk3hgcfqcwtf8j8yss6h6elqdeaerr8w89062u7v220tuqyqun8v": {AdaAmount: 1000000},
		"addr1qysfnqjx7rpg2l8g8ynjefwly5uftu5anqj3ze6939ywewm9nlxjsyf63x8tkxjagd70lea98mawg80k8lvjv2zma0fsq7xzls": {AdaAmount: 1000000},
		"addr1qx3g0pnxmmggj7jxckf064wzlfqjnyn9n4a2g7hn23wnjwcyzrx8ppcedh89pmnxdjm2hl4730eu38vzmfahxf22tgxs5hq8ja": {AdaAmount: 7555538},
		"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 5460362},
		"addr1qxlqxcxxn88fnf37wtgq8jwq8mfzumhq7svsnnr9crh0gkp0pxyg33eksp7m0p4atqcye7e3fzlgemnmwhcp26lwk4lq0fl74s": {AdaAmount: 1422183},
		"addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l": {AdaAmount: 1146843},
		"addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z": {AdaAmount: 1000000},
	}

	singlePaymentAddress := map[string]PaymentDetails{
		"addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah": {AdaAmount: 1253293},
	}

	It("should produce tx output correctly with many payment addresses and no source tokens", func() {
		aUTXO := generateTestUTXO(map[string]int{})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100205435/tx.raw"
		currentSlot := 48240738
		returnTxOut := 193402911
		minFee := 239685

		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, manyPaymentAddresses, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponseManyPaymentsNoSendTokens))
	})

	It("should produce tx output correctly with many payment addresses and some source tokens", func() {
		aUTXO := generateTestUTXO(map[string]int{
			"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 1,
			"1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY":        1,
		})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100205435/tx.raw"
		currentSlot := 48240738
		returnTxOut := 193402911
		minFee := 239685
		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, manyPaymentAddresses, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponseManyPaymentsNoSendTokensButPaymentAddrHasSourceTokens))
	})

	It("should produce tx output correctly with many payment addresses and one source token", func() {
		expectedResponse := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
			"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
			"--tx-out", "addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk+3108141",
			"--tx-out", "addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl+1000000",
			"--tx-out", "addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd+5460362",
			"--tx-out", "addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su+74355476",
			"--tx-out", "addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l+1146843",
			"--tx-out", "addr1qx3g0pnxmmggj7jxckf064wzlfqjnyn9n4a2g7hn23wnjwcyzrx8ppcedh89pmnxdjm2hl4730eu38vzmfahxf22tgxs5hq8ja+7555538",
			"--tx-out", "addr1qxdkv4x590q00jtkd2pvhg5uejfgc29pxafj6wvelkl5msrk3hgcfqcwtf8j8yss6h6elqdeaerr8w89062u7v220tuqyqun8v+1000000",
			"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
			"--tx-out", "addr1qxlqxcxxn88fnf37wtgq8jwq8mfzumhq7svsnnr9crh0gkp0pxyg33eksp7m0p4atqcye7e3fzlgemnmwhcp26lwk4lq0fl74s+1422183",
			"--tx-out", "addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl+3584514",
			"--tx-out", "addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z+1000000",
			"--tx-out", "addr1qysfnqjx7rpg2l8g8ynjefwly5uftu5anqj3ze6939ywewm9nlxjsyf63x8tkxjagd70lea98mawg80k8lvjv2zma0fsq7xzls+1000000",
			"--invalid-hereafter", "48250738", "--fee", "239685", "--out-file", "/tmp/pay_multi3100205435/tx.raw"}
		aUTXO := generateTestUTXO(map[string]int{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 1})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100205435/tx.raw"
		currentSlot := 48240738
		returnTxOut := 193402911
		minFee := 239685

		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, manyPaymentAddresses, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponse))
	})

	It("should produce tx output correctly with one payment address and some source tokens", func() {
		expectedResponse := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
			"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
			"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
			"--invalid-hereafter", "69250738", "--fee", "168685", "--out-file", "/tmp/pay_multi3100209999/tx.raw"}
		aUTXO := generateTestUTXO(map[string]int{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 1})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100209999/tx.raw"
		currentSlot := 69240738
		returnTxOut := 193402911
		minFee := 168685

		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, singlePaymentAddress, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponse))
	})

	It("should produce tx output correctly with one payment address and one source token", func() {
		expectedResponse := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
			"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
			"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
			"--invalid-hereafter", "69250738", "--fee", "168685", "--out-file", "/tmp/pay_multi3100209999/tx.raw"}
		aUTXO := generateTestUTXO(map[string]int{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 1})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100209999/tx.raw"
		currentSlot := 69240738
		returnTxOut := 193402911
		minFee := 168685
		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, singlePaymentAddress, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponse))
	})

	It("should produce tx output correctly with one payment address and many source tokens", func() {
		expectedResponse := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
			"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+1 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY+1 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
			"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293",
			"--invalid-hereafter", "69250738", "--fee", "168685", "--out-file", "/tmp/pay_multi3100209999/tx.raw"}
		aUTXO := generateTestUTXO(map[string]int{
			"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 1,
			"1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY":        1,
		})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100209999/tx.raw"
		currentSlot := 69240738
		returnTxOut := 193402911
		minFee := 168685
		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, singlePaymentAddress, []TokenDetails{})
		Expect(computedString).To(Equal(expectedResponse))
	})
})

var _ = Describe("Build transaction Payment With Tokens", func() {
	singlePaymentAddressWithOneToken := map[string]PaymentDetails{
		"addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah": {AdaAmount: 1253293, PaymentTokens: []TokenDetails{
			{TokenID: "2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons", TokenAmount: 37},
		}},
	}
	expectedResponseSinglePaymentWithOneToken := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
		"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+63 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons+50 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY",
		"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293+37 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons",
		"--invalid-hereafter", "85250738", "--fee", "178685", "--out-file", "/tmp/pay_multi3100208888/tx.raw"}

	singlePaymentAddressWithTwoTokens := map[string]PaymentDetails{
		"addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah": {AdaAmount: 1253293, PaymentTokens: []TokenDetails{
			{TokenID: "2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons", TokenAmount: 37},
			{TokenID: "45ace7db4aec426e119445e867816f31cdebc014b4f642fc1decda41.HONEYChristmas", TokenAmount: 42},
		}},
	}

	expectedResponseSinglePaymentWithTwoTokens := []string{"transaction", "build-raw", "--tx-in", "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80#0",
		"--tx-out", "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh+193402911+63 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons+38 45ace7db4aec426e119445e867816f31cdebc014b4f642fc1decda41.HONEYChristmas",
		"--tx-out", "addr1qxgtvpa3rvg3snl63vpjfs2s56fcyuksvhxhmaxldx0d9744l86svtmlgys63qgnggfnl6v9hwjflde37g9ys2wldy0q8ae6ah+1253293+37 2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons+42 45ace7db4aec426e119445e867816f31cdebc014b4f642fc1decda41.HONEYChristmas",
		"--invalid-hereafter", "85250738", "--fee", "178685", "--out-file", "/tmp/pay_multi3100208888/tx.raw"}

	It("should correctly build tx output when sending one payment with one token", func() {
		aUTXO := generateTestUTXO(map[string]int{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons": 100, "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY": 50})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100208888/tx.raw"
		currentSlot := 85240738
		returnTxOut := 193402911
		minFee := 178685
		paymentTokenDetails := []TokenDetails{{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons", 63}}
		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, singlePaymentAddressWithOneToken, paymentTokenDetails)
		Expect(computedString).To(Equal(expectedResponseSinglePaymentWithOneToken))
	})

	It("should correctly build tx output when sending one payment with two tokens", func() {
		aUTXO := generateTestUTXO(map[string]int{
			"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons":   100,
			"45ace7db4aec426e119445e867816f31cdebc014b4f642fc1decda41.HONEYChristmas": 80})
		sourceAddress := "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
		fileName := "/tmp/pay_multi3100208888/tx.raw"
		currentSlot := 85240738
		returnTxOut := 193402911
		minFee := 178685
		paymentTokenDetails := []TokenDetails{{"2f4157f71feaca0afb3122c21050ef82f8d39b6266075b17ba4a7b6a.TYHoskinsons", 63}, {"45ace7db4aec426e119445e867816f31cdebc014b4f642fc1decda41.HONEYChristmas", 38}}
		computedString := buildRawTransactionArgs(aUTXO, sourceAddress, fileName, currentSlot, returnTxOut, minFee, singlePaymentAddressWithTwoTokens, paymentTokenDetails)
		Expect(computedString).To(Equal(expectedResponseSinglePaymentWithTwoTokens))
	})
})

func generateTestUTXO(tokens map[string]int) *FullUTXO {
	return &FullUTXO{
		Address:       "",
		ADABalance:    999,
		TokenBalances: tokens, // Bit of a hack as we only have one row the balance is value of the row
		TXCount:       1,
		Rows: []UTXORow{
			{
				Hash:            "1e7fd3c8f66586035f64192c1c09d2438ee9ccac30f6fcd7215ac4d6656d7d80",
				TxID:            "0",
				LovelaceBalance: 1000,
				Tokens:          tokens,
			},
		},
	}
}
