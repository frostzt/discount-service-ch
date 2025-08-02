package testdata

import (
	"github.com/frostzt/discount/models"
	"github.com/shopspring/decimal"
)

func FakeCartData() []models.CartItem {
	return []models.CartItem{
		{
			Product: models.Product{
				ID:        "1",
				Brand:     "PUMA",
				BrandTier: models.BrandTierRegular,
				Category:  "T-shirts",
				BasePrice: decimal.NewFromFloat(1000),
			},
			Quantity: 1,
			Size:     "M",
		},
	}
}

func FakeCustomer() models.CustomerProfile {
	return models.CustomerProfile{
		ID:   "cust-123",
		Tier: "gold",
	}
}

func FakeSilverCustomer() models.CustomerProfile {
	return models.CustomerProfile{
		ID:   "cust-222",
		Tier: "silver",
	}
}

func FakeNikeCartItem() []models.CartItem {
	return []models.CartItem{
		{
			Product: models.Product{
				ID:        "2",
				Brand:     "Nike",
				BrandTier: models.BrandTierPremium,
				Category:  "T-shirts",
				BasePrice: decimal.NewFromFloat(2000),
			},
			Quantity: 1,
			Size:     "L",
		},
	}
}

func FakePaymentInfo() *models.PaymentInfo {
	bank := "ICICI"
	return &models.PaymentInfo{
		Method:   "CARD",
		BankName: &bank,
		CardType: nil,
	}
}
