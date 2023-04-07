package sendfunds

import (
	"bufio"
	"strconv"
	"strings"
)

type FullUTXO struct {
	Address       string
	ADABalance    int
	TokenBalances map[string]int
	TXCount       int
	Rows          []UTXORow
}

func NewFullUTXO() FullUTXO {
	tokenBalances := make(map[string]int)
	return FullUTXO{TokenBalances: tokenBalances}
}

func (f *FullUTXO) addRow(row UTXORow) {
	f.ADABalance += row.LovelaceBalance
	f.TXCount++
	for tokenID, tokenAmount := range row.Tokens {
		f.TokenBalances[tokenID] += tokenAmount
	}
	f.Rows = append(f.Rows, row)
}

func (f *FullUTXO) getTXInParam() []string {
	var stringParams []string
	for _, utxoRow := range f.Rows {
		stringParams = append(stringParams, "--tx-in")
		stringParams = append(stringParams, utxoRow.Hash+"#"+utxoRow.TxID)
	}
	return stringParams
}

type UTXORow struct {
	Hash            string
	TxID            string
	LovelaceBalance int
	Tokens          map[string]int
}

func shouldSkipRow(row string) bool {
	if strings.Contains(row, "TxHash") && strings.Contains(row, "TxIx") {
		return true
	}
	if strings.Contains(row, "-------------") {
		return true
	}
	if row == "" {
		return true
	}
	return false
}

func parseUTXORow(row string) (*UTXORow, error) {
	fields := strings.Fields(row)
	bal, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, err
	}
	tokens := make(map[string]int)
	tokenStrings := strings.Split(row, "+")
	if len(tokenStrings) <= 2 {
		return &UTXORow{fields[0], fields[1], bal, tokens}, nil
	}
	for _, tokenString := range tokenStrings[1 : len(tokenStrings)-1] {
		tokenFields := strings.Fields(strings.TrimSpace(tokenString))
		tokenBal, err := strconv.Atoi(tokenFields[0])
		if err != nil {
			return nil, err
		}
		tokens[tokenFields[1]] = tokenBal
	}
	return &UTXORow{fields[0], fields[1], bal, tokens}, nil
}

func parseFullUTXO(utxo, address string) (*FullUTXO, error) {
	newBalance := NewFullUTXO()
	newBalance.Address = address
	scanner := bufio.NewScanner(strings.NewReader(utxo))
	for scanner.Scan() {
		readRow := strings.TrimSpace(scanner.Text())
		if !shouldSkipRow(readRow) {
			utxoRow, err := parseUTXORow(readRow)
			if err != nil {
				return nil, err
			}
			newBalance.addRow(*utxoRow)
		}
	}
	return &newBalance, nil
}
