package models

import (
	"github.com/shopspring/decimal"
)

type DiscountedPrice struct {
	OriginalPrice    decimal.Decimal            `json:"original_price"`
	FinalPrice       decimal.Decimal            `json:"final_price"`
	AppliedDiscounts map[string]decimal.Decimal `json:"applied_discounts"`
	Message          string                     `json:"message"`
}
