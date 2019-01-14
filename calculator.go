package iban

import "fmt"

func Calculate(number, identityCode string) (string, error) {
	// https://www.cnb.cz/miranda2/export/sites/www.cnb.cz/en/payment_systems/accounts_bank_codes/download/bank_codes_CR_128.pdf

	switch identityCode {
	case
		"0100", // Komerční banka, a.s.
		"0300", // Československá obchodní banka, a.s.
		"2600", // Citibank Europe plc, organizační složka
		"3030", // Air Bank. a.s.
		"2700", // UniCredit Bank Czech Republic, a.s.
		"0600", // GE Money Bank, a.s.
		"0800", // Česká spořitelna, a.s.
		"5500", // Raiffeisenbank, a.s.
		"6210", // mBank, a.s
		"2010", // Fio, družstevní záložna
		"0710", // Česká národní banka
		"0730": // Česká národní banka - Clearing centre
		{
			return CalculateCzech(number, identityCode)
		}

	default:
		{
			return "", fmt.Errorf("Unsupported bankCode %s", identityCode)
		}
	}
}
