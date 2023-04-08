package sendfunds

import (
	"encoding/json"
	"fmt"
	"github.com/lambda-honeypot/ccli-tz/pkg/leader"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// FundSender some struct
type FundSender struct {
	network string
	magic   string
	runner  leader.CommandRunner
}

// NewFundSender init method for fundsender
func NewFundSender(runner leader.CommandRunner, network, magic string) FundSender {
	return FundSender{network, magic, runner}
}

// TipQuery struct to unmarshall json
type TipQuery struct {
	Slot int
}

func (fs *FundSender) RunSendFunds(startAddress, signingKeyFile string, paymentAddressesWithTokens map[string]PaymentDetails) error {
	balance, err := fs.createUTXOFromAddress(startAddress)
	if err != nil {
		return fmt.Errorf("failed to create UTXO from start address: %s with: %v", startAddress, err)
	}
	log.Infof("Balance Before: %d", balance.ADABalance)
	for idx, tokenBalance := range balance.TokenBalances {
		log.Debugf("Token Balance Before: %s + %d", idx, tokenBalance)
	}
	err = fs.payMultiple(startAddress, signingKeyFile, paymentAddressesWithTokens)
	if err != nil {
		return fmt.Errorf("failed to pay multiple wallets with: %v", err)
	}
	newBalance, err := fs.createUTXOFromAddress(startAddress)
	if err != nil {
		return fmt.Errorf("failed to create UTXO from start address: %s with: %v", startAddress, err)
	}
	log.Infof("Balance After: %+v\n", newBalance)
	return nil
}

func (fs *FundSender) createParamsFile(paramsFile string) error {
	queryProtocolArgs := fs.queryProtocolParamsArgs(paramsFile)
	log.Debugf("%s", queryProtocolArgs)
	queryProtocolReturn, err := fs.runner.RunCardanoCmd(queryProtocolArgs)
	log.Debugf("%s", string(queryProtocolReturn))
	if err != nil {
		return fmt.Errorf("stdin: %s stderr: %v", queryProtocolReturn, err)
	}
	return nil
}

func (fs *FundSender) sendTransaction(txSignedFile string) error {
	commandArgs := []string{"transaction", "submit", "--tx-file", txSignedFile, fs.network, fs.magic}
	txSubmitReturn, err := fs.runner.RunCardanoCmd(commandArgs)
	log.Infof("%s", string(txSubmitReturn))
	if err != nil {
		return fmt.Errorf("stdin: %s stderr: %v", txSubmitReturn, err)
	}
	return nil
}

func (fs *FundSender) signTransactionFile(txRawFile, signingKeyFile, txSignedFile string) error {
	commandArgs := []string{"transaction", "sign", "--tx-body-file", txRawFile, "--signing-key-file", signingKeyFile, fs.network, fs.magic, "--out-file", txSignedFile}
	txSignReturn, err := fs.runner.RunCardanoCmd(commandArgs)
	log.Infof("%s", txSignReturn)
	if err != nil {
		return fmt.Errorf("stdin: %s stderr: %v", txSignReturn, err)
	}
	return nil
}

func (fs *FundSender) getCurrentSlot() (int, error) {
	var tipQuery TipQuery
	commandArgs := []string{"query", "tip", fs.network, fs.magic}
	jsQuery, err := fs.runner.RunCardanoCmd(commandArgs)
	log.Debugf("Query tip: %s", jsQuery)
	if err != nil {
		return 0, fmt.Errorf("stdin: %s stderr: %v", jsQuery, err)
	}
	err = json.Unmarshal([]byte(jsQuery), &tipQuery)
	if err != nil {
		return 0, err
	}
	return tipQuery.Slot, nil
}

func (fs *FundSender) createUTXOFromAddress(tokenAddress string) (*FullUTXO, error) {
	commandArgs := []string{"query", "utxo", "--address", tokenAddress, fs.network, fs.magic}
	log.Debugf("calling cardano command with %v", commandArgs)
	queryReturn, err := fs.runner.RunCardanoCmd(commandArgs)
	if err != nil {
		return nil, fmt.Errorf("stdin: %s stderr: %v", queryReturn, err)
	}
	log.Debugf("UTXO query result\n%s", queryReturn)
	fullUTXO, err := parseFullUTXO(queryReturn, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to parse utxo query: %s", queryReturn)
	}

	return fullUTXO, nil
}

func (fs *FundSender) createRawTxFile(utxo *FullUTXO, sourceAddress, outFile string, paymentAddresses map[string]PaymentDetails, currentSlot, txOutAdaAmount, minFee int, txOutTokenAmounts []TokenDetails) error {
	rawTxArgs := buildRawTransactionArgs(utxo, sourceAddress, outFile, currentSlot, txOutAdaAmount, minFee, paymentAddresses, txOutTokenAmounts)
	log.Debugf("%s", rawTxArgs)
	buildRawReturn, err := fs.runner.RunCardanoCmd(rawTxArgs)
	log.Debugf("%s", string(buildRawReturn))
	if err != nil {
		return fmt.Errorf("stdin: %s stderr: %v", buildRawReturn, err)
	}
	return nil
}

func (fs *FundSender) calculateMinimumFee(utxo *FullUTXO, paymentAddresses map[string]PaymentDetails, tempFile, paramsFile string) (int, error) {
	transactionOutCount := len(paymentAddresses) + 1
	minFeeArgs := fs.calculateMinFeeArgs(paramsFile, tempFile, utxo.TXCount, transactionOutCount)
	log.Debugf("%s", minFeeArgs)
	minFeeReturn, err := fs.runner.RunCardanoCmd(minFeeArgs)
	log.Debugf("MIN FEE: %s", minFeeReturn)
	if err != nil {
		return 0, fmt.Errorf("stdin: %s stderr: %v", minFeeReturn, err)
	}
	minFeeSplit := strings.Fields(string(minFeeReturn))
	minFee, err := strconv.Atoi(minFeeSplit[0])
	if err != nil {
		return 0, err
	}
	return minFee, nil
}

func (fs *FundSender) payMultiple(sourceAddress, signingKeyFile string, paymentDetails map[string]PaymentDetails) error {
	dir, err := ioutil.TempDir("", "pay_multi")
	if err != nil {
		return err
	}
	tmpFile := dir + "/tx.tmp"
	rawFile := dir + "/tx.raw"
	txSignedFile := dir + "/tx.signed"
	paramsFile := dir + "/" + strings.ReplaceAll(fmt.Sprintf("%s-params.json", fs.network), "--", "")

	defer os.RemoveAll(dir)
	slot, _ := fs.getCurrentSlot()
	err = fs.createParamsFile(paramsFile)
	if err != nil {
		return fmt.Errorf("failed to create params file: %v", err)
	}
	utxoDetails, err := fs.createUTXOFromAddress(sourceAddress)
	if err != nil {
		return fmt.Errorf("failed to create utxofrom address: %s with error: %v", sourceAddress, err)
	}
	paymentTokenDetails, err := generateTokenDetailsAndVerify(utxoDetails, paymentDetails)
	if err != nil {
		return fmt.Errorf("failed to generate token details with: %v", err)
	}
	err = fs.createRawTxFile(utxoDetails, sourceAddress, tmpFile, paymentDetails, slot, 0, 0, []TokenDetails{})
	if err != nil {
		return fmt.Errorf("failed to create tmp tx file for fee calc: %v", err)
	}
	minFee, err := fs.calculateMinimumFee(utxoDetails, paymentDetails, tmpFile, paramsFile)
	if err != nil {
		return fmt.Errorf("failed to calculate min fee: %v", err)
	}
	log.Infof("calculate min fee: %d", minFee)
	totalADAinLovelace := 0
	for _, paymentDetail := range paymentDetails {
		totalADAinLovelace += paymentDetail.AdaAmount
	}
	txOutAdaAmount := utxoDetails.ADABalance - totalADAinLovelace - minFee
	err = fs.createRawTxFile(utxoDetails, sourceAddress, rawFile, paymentDetails, slot, txOutAdaAmount, minFee, paymentTokenDetails)
	if err != nil {
		return fmt.Errorf("failed to create raw tx file for payment: %v", err)
	}
	err = fs.signTransactionFile(rawFile, signingKeyFile, txSignedFile)
	if err != nil {
		return fmt.Errorf("failed to sign tx file for send: %v", err)
	}
	err = fs.sendTransaction(txSignedFile)
	if err != nil {
		return fmt.Errorf("failed to send signed tx: %v", err)
	}
	return nil
}

func generateTokenDetailsAndVerify(utxo *FullUTXO, paymentDetails map[string]PaymentDetails) ([]TokenDetails, error) {
	sendTotals := make(map[string]int)
	var returnTokens []TokenDetails
	for _, paymentDetail := range paymentDetails {
		for _, tokenDetail := range paymentDetail.PaymentTokens {
			sendTotals[tokenDetail.TokenID] += tokenDetail.TokenAmount
		}
	}
	for tokenID, sendTokenAmount := range sendTotals {
		if utxo.TokenBalances[tokenID] < sendTokenAmount {
			return nil, fmt.Errorf("total send token amount for token: %s is %d - this is greater than source wallet balance of %d", tokenID, sendTokenAmount, utxo.TokenBalances[tokenID])
		}
		adjustedAmount := utxo.TokenBalances[tokenID] - sendTokenAmount
		returnTokens = append(returnTokens, TokenDetails{TokenID: tokenID, TokenAmount: adjustedAmount})
	}
	return returnTokens, nil
}
