package models

type PaymentInfo struct {
	Method   string  `json:"method"` // CARD, UPI
	BankName *string `json:"bank_name"`
	CardType *string `json:"card_type"` // CREDIT, DEBIT
}

func (p *PaymentInfo) BankNameOrDefault() string {
	if p.BankName != nil {
		return *p.BankName
	}
	return ""
}
