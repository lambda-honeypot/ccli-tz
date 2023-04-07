package sendfunds

import (
	"github.com/lambda-honeypot/ccli-tz/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type PaymentYaml struct {
	SourceAddress   string                    `yaml:"sourceAddress"`
	TargetAddresses map[string]PaymentDetails `yaml:"targetAddresses"`
}

func ReadPaymentFile(fileReader utils.FileUtilsInterface, paymentFilePath string) *PaymentYaml {
	paymentFile, err := fileReader.ReadFile(paymentFilePath)
	if err != nil {
		log.Errorf("failed to read file contents of: %s with err: %v", paymentFile, err)
		return nil
	}
	paymentYaml := &PaymentYaml{}
	err = yaml.Unmarshal(paymentFile, paymentYaml)
	if err != nil {
		log.Errorf("failed to read yaml in file: %s with err: %v", paymentFile, err)
		return nil
	}
	return paymentYaml
}
