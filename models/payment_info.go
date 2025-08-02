package models

type PaymentInfo struct {
	Method   string  `json:"method"`
	BankName *string `json:"bank_name"`
	CardType *string `json:"card_type"`
}

func (p *PaymentInfo) BankNameOrDefault() string {
	if p.BankName != nil {
		return *p.BankName
	}
	return ""
}
