package sendfunds

import (
	"github.com/golang/mock/gomock"
	mock_utils "github.com/lambda-honeypot/ccli-tz/pkg/utils/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadPaymentFile", func() {
	var (
		ctrl           *gomock.Controller
		mockFileReader *mock_utils.MockFileUtilsInterface
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFileReader = mock_utils.NewMockFileUtilsInterface(ctrl)
	})
	paymentAddressesWithTokens := `sourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
targetAddresses:
  addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd:
    lovelaceAmount: 3899796
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv:
    lovelaceAmount: 2996535
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j:
    lovelaceAmount: 1871848
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q8dkapy5tr3tda6unwd2jvg65ym5gx2m4nn3slm7npzwnqfxv0ezkxrz50apjryr6atcn22myqeev8z5lrvxd4t46gzq9g86wq:
    lovelaceAmount: 1037712
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qxyfr6vuy50dxfltaszcqz2w8n7qdflns5rm0asjcehwvyf58pwmrh62tsrmjun5kz4qnxkzma50sv25q64hwtegy7fsudl76u:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q9kvaxpvyr9k7pvgqhsuqteh28pfv3z2fa85tcugayh2y9pknhhhacfn6rg5s38c2vfh5n9wp9qmfena6pm4ve2lty9qfnkx24:
    lovelaceAmount: 2328534
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qyy3y635lg2euthsu6l072njmrpshtwcdmsx6smfuvadzlzq6s7d9rhfy0v2nhg92wtztc98360l7tk2pp3qr0am4kts39pvn3:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qy5auf3n989lyk3tgy4cdtlxrznd03cjgwc8wfha50v7952w2x2wjpx7spu5yhx0p6ygcr0w29fvpvm8h6cp62qd6tdqsyn59q:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qxdfpflavnv765h79rh868a4xzxt05nw3pe036yrr0rge7tzj4ekqsfm4xmj7277v6ka8a33aqdg6xsc009gjy54ltcshfvee3:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qyppjer4dq62jknr5ew9uqehy4vjnryv2v97s7sh6d8kl5nrga6hs0dwjpx47efmgjh7lvdykkjp8tx26rvss0u8wqmqx3r4d2:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qymqa6wefry3dua7nhk842de9tf44qa70eu3fmay2nxe0wmdza0mmlvlpkkprx4en2lq9sqt9cwfd0f0fvnt3enn3wqqwde792:
    lovelaceAmount: 1000000
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q97uarqcraa40276pzysgjveh3fj6n2720m93ax78p6knqyf3ljhzkf0rpqwh2spu4ug9r8hm4z9azrj9hjyg6fanzysm7s0m5:
    lovelaceAmount: 14528316
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qxkwkdx4a5c02tc475ff4pf4g8fy7xc03fywh9lfjjeafn9s6yqler9lrl3mkc7ys7ltkd9y9wfjuwy5ve2f8egwj79q9ykz0v:
    lovelaceAmount: 4580949
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl:
    lovelaceAmount: 2735559
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1q85hhneggnv06g7qehjpfqehdzckmew3g0gmd3uw9s06r37xyn0jhe5rkad0yazxmpld060pucac3gfwp29mwj4y840s3gl7yt:
    lovelaceAmount: 1014211
    paymentTokens:
    - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
      tokenAmount: 1
  addr1qxk94n4knlan8q8xxate8rqyzx68j5pf2j20dxega5ys86kxn4ezgx2n9su07vqv2y5zm4canfhzx45z9d2ven64rmqq33jymw:
     lovelaceAmount: 2099163
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su:
     lovelaceAmount: 38884333
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk:
     lovelaceAmount: 2224737
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l:
     lovelaceAmount: 1000000
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1qx7qzfdjrlep38k8ddz2q50a33mmzssgtfgvn884rlkaf0lq5sam04dt7kklp76tkwhec88dyf7rs77qc69l8f4rmh4syy3twn:
     lovelaceAmount: 4421620
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1q90vwqsc09lqvqdwle287qvz3havcdpscmfpm5yf8hu2w5hfzp0szq0lv8jvem4epan4f8xncncwsjhhhu95pa5nmedsvphrjc:
     lovelaceAmount: 1852709
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1qxvn467qy6amn6shwxrapmxt899pdk7pedwtz40mm39ql600ymlszes53w6dd39e2ygfflk5yjz2y59hemlp7j608wrs0dram2:
     lovelaceAmount: 9539021
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
  addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z:
     lovelaceAmount: 1000000
     paymentTokens:
     - tokenID: 1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY
       tokenAmount: 1
`
	paymentAddressesNoTokens := `sourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
targetAddresses:
  addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd: 
    lovelaceAmount: 3899796
  addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs: 
    lovelaceAmount: 1000000
  addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv: 
    lovelaceAmount: 2996535
  addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j: 
    lovelaceAmount: 1871848
  addr1q8dkapy5tr3tda6unwd2jvg65ym5gx2m4nn3slm7npzwnqfxv0ezkxrz50apjryr6atcn22myqeev8z5lrvxd4t46gzq9g86wq: 
    lovelaceAmount: 1037712
  addr1qxyfr6vuy50dxfltaszcqz2w8n7qdflns5rm0asjcehwvyf58pwmrh62tsrmjun5kz4qnxkzma50sv25q64hwtegy7fsudl76u: 
    lovelaceAmount: 1000000
  addr1q9kvaxpvyr9k7pvgqhsuqteh28pfv3z2fa85tcugayh2y9pknhhhacfn6rg5s38c2vfh5n9wp9qmfena6pm4ve2lty9qfnkx24: 
    lovelaceAmount: 2328534
  addr1qyy3y635lg2euthsu6l072njmrpshtwcdmsx6smfuvadzlzq6s7d9rhfy0v2nhg92wtztc98360l7tk2pp3qr0am4kts39pvn3: 
    lovelaceAmount: 1000000
  addr1qy5auf3n989lyk3tgy4cdtlxrznd03cjgwc8wfha50v7952w2x2wjpx7spu5yhx0p6ygcr0w29fvpvm8h6cp62qd6tdqsyn59q: 
    lovelaceAmount: 1000000
  addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl: 
    lovelaceAmount: 1000000
  addr1qxdfpflavnv765h79rh868a4xzxt05nw3pe036yrr0rge7tzj4ekqsfm4xmj7277v6ka8a33aqdg6xsc009gjy54ltcshfvee3: 
    lovelaceAmount: 1000000
  addr1qyppjer4dq62jknr5ew9uqehy4vjnryv2v97s7sh6d8kl5nrga6hs0dwjpx47efmgjh7lvdykkjp8tx26rvss0u8wqmqx3r4d2: 
    lovelaceAmount: 1000000
  addr1qymqa6wefry3dua7nhk842de9tf44qa70eu3fmay2nxe0wmdza0mmlvlpkkprx4en2lq9sqt9cwfd0f0fvnt3enn3wqqwde792: 
    lovelaceAmount: 1000000
  addr1q97uarqcraa40276pzysgjveh3fj6n2720m93ax78p6knqyf3ljhzkf0rpqwh2spu4ug9r8hm4z9azrj9hjyg6fanzysm7s0m5: 
    lovelaceAmount: 14528316
  addr1qxkwkdx4a5c02tc475ff4pf4g8fy7xc03fywh9lfjjeafn9s6yqler9lrl3mkc7ys7ltkd9y9wfjuwy5ve2f8egwj79q9ykz0v: 
    lovelaceAmount: 4580949
  addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl: 
    lovelaceAmount: 2735559
  addr1q85hhneggnv06g7qehjpfqehdzckmew3g0gmd3uw9s06r37xyn0jhe5rkad0yazxmpld060pucac3gfwp29mwj4y840s3gl7yt: 
    lovelaceAmount: 1014211
  addr1qxk94n4knlan8q8xxate8rqyzx68j5pf2j20dxega5ys86kxn4ezgx2n9su07vqv2y5zm4canfhzx45z9d2ven64rmqq33jymw: 
    lovelaceAmount: 2099163
  addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su: 
    lovelaceAmount: 38884333
  addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk: 
    lovelaceAmount: 2224737
  addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l: 
    lovelaceAmount: 1000000
  addr1qx7qzfdjrlep38k8ddz2q50a33mmzssgtfgvn884rlkaf0lq5sam04dt7kklp76tkwhec88dyf7rs77qc69l8f4rmh4syy3twn: 
    lovelaceAmount: 4421620
  addr1q90vwqsc09lqvqdwle287qvz3havcdpscmfpm5yf8hu2w5hfzp0szq0lv8jvem4epan4f8xncncwsjhhhu95pa5nmedsvphrjc: 
    lovelaceAmount: 1852709
  addr1qxvn467qy6amn6shwxrapmxt899pdk7pedwtz40mm39ql600ymlszes53w6dd39e2ygfflk5yjz2y59hemlp7j608wrs0dram2: 
    lovelaceAmount: 9539021
  addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z: 
    lovelaceAmount: 1000000
`
	It("should read the payment file with tokens as expected", func() {
		paymentFilePath := "/some/fictitious/path.yml"
		mockFileReader.EXPECT().ReadFile(paymentFilePath).Return([]byte(paymentAddressesWithTokens), nil)
		paymentYaml := ReadPaymentFile(mockFileReader, paymentFilePath)
		expectedPaymentYaml := PaymentYaml{
			SourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh",
			TargetAddresses: map[string]PaymentDetails{
				"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 3899796,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv": {AdaAmount: 2996535,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j": {AdaAmount: 1871848,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q8dkapy5tr3tda6unwd2jvg65ym5gx2m4nn3slm7npzwnqfxv0ezkxrz50apjryr6atcn22myqeev8z5lrvxd4t46gzq9g86wq": {AdaAmount: 1037712,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qxyfr6vuy50dxfltaszcqz2w8n7qdflns5rm0asjcehwvyf58pwmrh62tsrmjun5kz4qnxkzma50sv25q64hwtegy7fsudl76u": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q9kvaxpvyr9k7pvgqhsuqteh28pfv3z2fa85tcugayh2y9pknhhhacfn6rg5s38c2vfh5n9wp9qmfena6pm4ve2lty9qfnkx24": {AdaAmount: 2328534,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qyy3y635lg2euthsu6l072njmrpshtwcdmsx6smfuvadzlzq6s7d9rhfy0v2nhg92wtztc98360l7tk2pp3qr0am4kts39pvn3": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qy5auf3n989lyk3tgy4cdtlxrznd03cjgwc8wfha50v7952w2x2wjpx7spu5yhx0p6ygcr0w29fvpvm8h6cp62qd6tdqsyn59q": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qxdfpflavnv765h79rh868a4xzxt05nw3pe036yrr0rge7tzj4ekqsfm4xmj7277v6ka8a33aqdg6xsc009gjy54ltcshfvee3": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qyppjer4dq62jknr5ew9uqehy4vjnryv2v97s7sh6d8kl5nrga6hs0dwjpx47efmgjh7lvdykkjp8tx26rvss0u8wqmqx3r4d2": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qymqa6wefry3dua7nhk842de9tf44qa70eu3fmay2nxe0wmdza0mmlvlpkkprx4en2lq9sqt9cwfd0f0fvnt3enn3wqqwde792": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q97uarqcraa40276pzysgjveh3fj6n2720m93ax78p6knqyf3ljhzkf0rpqwh2spu4ug9r8hm4z9azrj9hjyg6fanzysm7s0m5": {AdaAmount: 14528316,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qxkwkdx4a5c02tc475ff4pf4g8fy7xc03fywh9lfjjeafn9s6yqler9lrl3mkc7ys7ltkd9y9wfjuwy5ve2f8egwj79q9ykz0v": {AdaAmount: 4580949,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 2735559,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q85hhneggnv06g7qehjpfqehdzckmew3g0gmd3uw9s06r37xyn0jhe5rkad0yazxmpld060pucac3gfwp29mwj4y840s3gl7yt": {AdaAmount: 1014211,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qxk94n4knlan8q8xxate8rqyzx68j5pf2j20dxega5ys86kxn4ezgx2n9su07vqv2y5zm4canfhzx45z9d2ven64rmqq33jymw": {AdaAmount: 2099163,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su": {AdaAmount: 38884333,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk": {AdaAmount: 2224737,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qx7qzfdjrlep38k8ddz2q50a33mmzssgtfgvn884rlkaf0lq5sam04dt7kklp76tkwhec88dyf7rs77qc69l8f4rmh4syy3twn": {AdaAmount: 4421620,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1q90vwqsc09lqvqdwle287qvz3havcdpscmfpm5yf8hu2w5hfzp0szq0lv8jvem4epan4f8xncncwsjhhhu95pa5nmedsvphrjc": {AdaAmount: 1852709,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qxvn467qy6amn6shwxrapmxt899pdk7pedwtz40mm39ql600ymlszes53w6dd39e2ygfflk5yjz2y59hemlp7j608wrs0dram2": {AdaAmount: 9539021,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
				"addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z": {AdaAmount: 1000000,
					PaymentTokens: []TokenDetails{{TokenID: "1815bee29d9d1eabf78b7f21f29ae55cbad8d06fa470a65ddbf98156.HONEY", TokenAmount: 1}}},
			},
		}
		Expect(paymentYaml).ToNot(BeNil())
		Expect(paymentYaml.SourceAddress).To(Equal(expectedPaymentYaml.SourceAddress))
		for k, v := range expectedPaymentYaml.TargetAddresses {
			testValue := paymentYaml.TargetAddresses[k]
			Expect(testValue).To(Not(BeNil()))
			Expect(testValue).To(Equal(v))
		}
	})
	It("should read the payment file without tokens as expected", func() {
		paymentFilePath := "/some/fictitious/path.yml"
		mockFileReader.EXPECT().ReadFile(paymentFilePath).Return([]byte(paymentAddressesNoTokens), nil)
		paymentYaml := ReadPaymentFile(mockFileReader, paymentFilePath)
		expectedPaymentYaml := PaymentYaml{
			SourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh",
			TargetAddresses: map[string]PaymentDetails{
				"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 3899796},
				"addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs": {AdaAmount: 1000000},
				"addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv": {AdaAmount: 2996535},
				"addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j": {AdaAmount: 1871848},
				"addr1q8dkapy5tr3tda6unwd2jvg65ym5gx2m4nn3slm7npzwnqfxv0ezkxrz50apjryr6atcn22myqeev8z5lrvxd4t46gzq9g86wq": {AdaAmount: 1037712},
				"addr1qxyfr6vuy50dxfltaszcqz2w8n7qdflns5rm0asjcehwvyf58pwmrh62tsrmjun5kz4qnxkzma50sv25q64hwtegy7fsudl76u": {AdaAmount: 1000000},
				"addr1q9kvaxpvyr9k7pvgqhsuqteh28pfv3z2fa85tcugayh2y9pknhhhacfn6rg5s38c2vfh5n9wp9qmfena6pm4ve2lty9qfnkx24": {AdaAmount: 2328534},
				"addr1qyy3y635lg2euthsu6l072njmrpshtwcdmsx6smfuvadzlzq6s7d9rhfy0v2nhg92wtztc98360l7tk2pp3qr0am4kts39pvn3": {AdaAmount: 1000000},
				"addr1qy5auf3n989lyk3tgy4cdtlxrznd03cjgwc8wfha50v7952w2x2wjpx7spu5yhx0p6ygcr0w29fvpvm8h6cp62qd6tdqsyn59q": {AdaAmount: 1000000},
				"addr1q8k63n24xq4qkzh8skjdjjuph5eywggkdhk6sv06qscwd7zh5y07rjxpf9r407yuu0v9u98c8hfk0462y6hck3fyyaxs24xupl": {AdaAmount: 1000000},
				"addr1qxdfpflavnv765h79rh868a4xzxt05nw3pe036yrr0rge7tzj4ekqsfm4xmj7277v6ka8a33aqdg6xsc009gjy54ltcshfvee3": {AdaAmount: 1000000},
				"addr1qyppjer4dq62jknr5ew9uqehy4vjnryv2v97s7sh6d8kl5nrga6hs0dwjpx47efmgjh7lvdykkjp8tx26rvss0u8wqmqx3r4d2": {AdaAmount: 1000000},
				"addr1qymqa6wefry3dua7nhk842de9tf44qa70eu3fmay2nxe0wmdza0mmlvlpkkprx4en2lq9sqt9cwfd0f0fvnt3enn3wqqwde792": {AdaAmount: 1000000},
				"addr1q97uarqcraa40276pzysgjveh3fj6n2720m93ax78p6knqyf3ljhzkf0rpqwh2spu4ug9r8hm4z9azrj9hjyg6fanzysm7s0m5": {AdaAmount: 14528316},
				"addr1qxkwkdx4a5c02tc475ff4pf4g8fy7xc03fywh9lfjjeafn9s6yqler9lrl3mkc7ys7ltkd9y9wfjuwy5ve2f8egwj79q9ykz0v": {AdaAmount: 4580949},
				"addr1qy5wy8wkl4kkgzstdke8wgnz6q0qfxr3dmp4whf09f59nz9jh2hy3lgnjzf8gtqhu34hcv0g7xc2v9qdl63c2qvqzuts2r9ryl": {AdaAmount: 2735559},
				"addr1q85hhneggnv06g7qehjpfqehdzckmew3g0gmd3uw9s06r37xyn0jhe5rkad0yazxmpld060pucac3gfwp29mwj4y840s3gl7yt": {AdaAmount: 1014211},
				"addr1qxk94n4knlan8q8xxate8rqyzx68j5pf2j20dxega5ys86kxn4ezgx2n9su07vqv2y5zm4canfhzx45z9d2ven64rmqq33jymw": {AdaAmount: 2099163},
				"addr1q9l52r7ta9zp957zapxlrsvmmvfpka25ltd27zcrt2caeg7458rnmjp9dgr3kz79terqcwkae9yhk9uue8fchn3pnvhq3yx2su": {AdaAmount: 38884333},
				"addr1q8du6tsm3qr9dpq53vqtrc983ssjshvhk25dzce42hnvp7778dkswup2zt9hcanyar93njuj0xnl8qtwtmw8cezz3xwshzstvk": {AdaAmount: 2224737},
				"addr1qx0rjvslgv2j6e5s9kmcyl79v6h7yksxdphkwy9274j2ldhqx06zs3s9lt43rrff7zg4zctfwm8s6v00djh4k4q2tlmqr3v24l": {AdaAmount: 1000000},
				"addr1qx7qzfdjrlep38k8ddz2q50a33mmzssgtfgvn884rlkaf0lq5sam04dt7kklp76tkwhec88dyf7rs77qc69l8f4rmh4syy3twn": {AdaAmount: 4421620},
				"addr1q90vwqsc09lqvqdwle287qvz3havcdpscmfpm5yf8hu2w5hfzp0szq0lv8jvem4epan4f8xncncwsjhhhu95pa5nmedsvphrjc": {AdaAmount: 1852709},
				"addr1qxvn467qy6amn6shwxrapmxt899pdk7pedwtz40mm39ql600ymlszes53w6dd39e2ygfflk5yjz2y59hemlp7j608wrs0dram2": {AdaAmount: 9539021},
				"addr1qy952ujwhtead56xsrwycjqwsr7hzsy6p8jgqrraxhnmj3h6tqesle4sxk6vuhem9sw9843wp48tsyv9xh2gkvch7qdsx8v75z": {AdaAmount: 1000000},
			},
		}
		Expect(paymentYaml).ToNot(BeNil())
		Expect(paymentYaml.SourceAddress).To(Equal(expectedPaymentYaml.SourceAddress))
		for k, v := range expectedPaymentYaml.TargetAddresses {
			testValue := paymentYaml.TargetAddresses[k]
			Expect(testValue).To(Not(BeNil()))
			Expect(testValue).To(Equal(v))
		}
	})
	It("should read the payment file without tokens as expected for smaller case", func() {
		paymentFilePath := "/some/fictitious/path.yml"
		paymentAddressesShore := `sourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh"
targetAddresses:
  addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd: 
    lovelaceAmount: 3899796
  addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs: 
    lovelaceAmount: 1000000
  addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv: 
    lovelaceAmount: 2996535
  addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j: 
    lovelaceAmount: 1871848`
		mockFileReader.EXPECT().ReadFile(paymentFilePath).Return([]byte(paymentAddressesShore), nil)
		paymentYaml := ReadPaymentFile(mockFileReader, paymentFilePath)
		expectedPaymentYaml := PaymentYaml{
			SourceAddress: "addr1q8q566cvhawynjmw008u5xlzkqaplx33vjhs82ec7f2vzt7m9dtqxjj5kv4u40r5ss7dsy679zcw9xkm07kasdg6u4hs3azrhh",
			TargetAddresses: map[string]PaymentDetails{
				"addr1q9fdv824xyfta60g47rq8pjy0ptsuexjlxjg2fvecw40jcsxqxwwhd4zts3nk07272nul2uyk7wgvgpyw3thdktmtpqq6427kd": {AdaAmount: 3899796},
				"addr1q89fmjw6qsk9924rle7qtp0fthnzsj7nrwd5zytt8pd68jq5tjvdt26nth9jn9utj7nwwcatkhstgv87mzyecrdl5zssvur6cs": {AdaAmount: 1000000},
				"addr1q9ufyayaepyxdvu3qhyxuq2p2uarrlay29ce35fwj9el22glwzh2pzj4yknexke487hkpe34s6l66rqtwkzy9asjgm5sx3c3jv": {AdaAmount: 2996535},
				"addr1q80830etk05d47du5frw5u3229hldy52ak5fr9qqvqzdzl3qc6t47l93dcx5ffw3czteaugj5za86duca9fckpyusy7s4fyv2j": {AdaAmount: 1871848},
			},
		}
		Expect(paymentYaml).ToNot(BeNil())
		Expect(paymentYaml.SourceAddress).To(Equal(expectedPaymentYaml.SourceAddress))
		for k, v := range expectedPaymentYaml.TargetAddresses {
			testValue := paymentYaml.TargetAddresses[k]
			Expect(testValue).To(Not(BeNil()))
			Expect(testValue).To(Equal(v))
		}
	})
})
