package sendfunds

import (
	"fmt"
	"sort"
)

type PaymentDetails struct {
	AdaAmount     int            `yaml:"lovelaceAmount"`
	PaymentTokens []TokenDetails `yaml:"paymentTokens"`
}

type TokenDetails struct {
	TokenID     string `yaml:"tokenID"`
	TokenAmount int    `yaml:"tokenAmount"`
}

func buildRawTransactionArgs(utxo *FullUTXO, sourceAddress, outFile string, currentSlot, txOutAdaAmount, minFee int, paymentAddresses map[string]PaymentDetails, txTokenOutAmounts []TokenDetails) []string {
	rawTxArgs := []string{"transaction", "build-raw"}
	rawTxArgs = append(rawTxArgs, utxo.getTXInParam()...)
	rawTxArgs = append(rawTxArgs, "--tx-out", deriveTxOutForSourceAddress(sourceAddress, txOutAdaAmount, txTokenOutAmounts, utxo))

	sortedKeys := make([]string, 0, len(paymentAddresses))
	for k := range paymentAddresses {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	for _, addr := range sortedKeys {
		paymentDetails := paymentAddresses[addr]
		rawTxArgs = append(rawTxArgs, "--tx-out")
		txStr := deriveTxOutForDestinationAddress(addr, paymentDetails)
		rawTxArgs = append(rawTxArgs, txStr)
	}

	rawTxArgs = append(rawTxArgs, "--invalid-hereafter")
	rawTxArgs = append(rawTxArgs, fmt.Sprintf("%d", currentSlot+10000))
	rawTxArgs = append(rawTxArgs, "--fee")
	rawTxArgs = append(rawTxArgs, fmt.Sprint(minFee))
	rawTxArgs = append(rawTxArgs, "--out-file")
	rawTxArgs = append(rawTxArgs, outFile)
	return rawTxArgs
}

func deriveTxOutForDestinationAddress(addr string, paymentDetails PaymentDetails) string {
	baseString := fmt.Sprintf("%s+%d", addr, paymentDetails.AdaAmount)
	for _, tokenDetails := range paymentDetails.PaymentTokens {
		baseString += fmt.Sprintf("+%d %s", tokenDetails.TokenAmount, tokenDetails.TokenID)
	}
	return baseString
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func deriveTxOutForSourceAddress(sourceAddress string, txOutAdaAmount int, txTokenOutAmounts []TokenDetails, utxo *FullUTXO) string {
	var tokensInPayment []string
	baseString := fmt.Sprintf("%s+%s", sourceAddress, fmt.Sprint(txOutAdaAmount))
	for _, tokenDetails := range txTokenOutAmounts {
		baseString += fmt.Sprintf("+%d %s", tokenDetails.TokenAmount, tokenDetails.TokenID)
		tokensInPayment = append(tokensInPayment, tokenDetails.TokenID)
	}

	sortedKeys := make([]string, 0, len(utxo.TokenBalances))
	for k := range utxo.TokenBalances {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	for _, tokenHash := range sortedKeys {
		if !stringInSlice(tokenHash, tokensInPayment) {
			tokenAmount := utxo.TokenBalances[tokenHash]
			baseString += fmt.Sprintf("+%d %s", tokenAmount, tokenHash)
		}
	}
	return baseString
}

func (fs *FundSender) calculateMinFeeArgs(paramsFile, tempFile string, transactionInCount, transactionOutCount int) []string {
	rawTxArgs := []string{"transaction", "calculate-min-fee", "--tx-body-file", tempFile, "--tx-in-count",
		fmt.Sprint(transactionInCount), "--tx-out-count", fmt.Sprint(transactionOutCount), fs.network, fs.magic, "--witness-count", "1",
		"--byron-witness-count", "0", "--protocol-params-file", paramsFile}
	return rawTxArgs
}

func (fs *FundSender) queryProtocolParamsArgs(paramsFile string) []string {
	rawTxArgs := []string{"query", "protocol-parameters", "--out-file", paramsFile, fs.network, fs.magic}
	return rawTxArgs
}
