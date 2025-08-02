package models

import (
	"github.com/shopspring/decimal"
)

type BrandTier string

const (
	BrandTierPremium BrandTier = "premium"
	BrandTierRegular BrandTier = "regular"
	BrandTierBudget  BrandTier = "budget"
)

type Product struct {
	ID           string          `json:"id"`
	Brand        string          `json:"brand"`
	BrandTier    BrandTier       `json:"brand_tier"`
	Category     string          `json:"category"`
	BasePrice    decimal.Decimal `json:"base_price"`
	CurrentPrice decimal.Decimal `json:"current_price"`
}

type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Size     string  `json:"size"`
}

type PaymentInfo struct {
	Method   string  `json:"method"`
	BankName *string `json:"bank_name"`
	CardType *string `json:"card_type"`
}

type DiscountedPrice struct {
	OriginalPrice    decimal.Decimal            `json:"original_price"`
	FinalPrice       decimal.Decimal            `json:"final_price"`
	AppliedDiscounts map[string]decimal.Decimal `json:"applied_discounts"`
	Message          string                     `json:"message"`
}

type CustomerProfile struct {
	ID   string `json:"id"`
	Tier string `json:"tier"`
}
