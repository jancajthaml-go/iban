package iban

import (
	"fmt"
	"strconv"
	"strings"
)

var checkSumToString = make([]string, 98)

func init() {
	for i := 0; i < 10; i++ {
		checkSumToString[i] = "0" + strconv.Itoa(i)
	}

	for i := 10; i < 98; i++ {
		checkSumToString[i] = strconv.Itoa(i)
	}
}

func calculateCzech(number, bankCode string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = ""
			err = fmt.Errorf("%+v", r)
		}
	}()

	// canonise input
	canonisedNumber := strings.Replace(number, "-", "", -1)
	// accountNumber of length 16
	paddedNumber := "0000000000000000"[0:16-len(canonisedNumber)] + canonisedNumber
	// bankCode of length 4
	paddedBankCode := "0000"[0:4-len(bankCode)] + bankCode
	// country code for "Czech Republic"
	countryCode := "CZ"
	// country code converted to digits
	countryDigits := "123500"
	// checksum mod 97
	checksum := (98 - mod97(paddedBankCode+paddedNumber+countryDigits))
	if checksum == 99 {
		return
	}

	result = countryCode + checkSumToString[checksum] + paddedBankCode + paddedNumber

	return
}
