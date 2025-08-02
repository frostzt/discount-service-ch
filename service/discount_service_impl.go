package service

import (
	"context"
	"errors"
	"strings"

	"github.com/frostzt/discount/models"
	"github.com/shopspring/decimal"
)

type discountService struct {
	validCodes map[string]struct{}
}

func NewDiscountService() DiscountService {
	return &discountService{
		validCodes: map[string]struct{}{
			"SUPER69": {},
		},
	}
}

func (d *discountService) CalculateCartDiscounts(ctx context.Context,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
	paymentInfo *models.PaymentInfo) (*models.DiscountedPrice, error) {

	var originalPrice decimal.Decimal
	var finalPrice decimal.Decimal
	applied := make(map[string]decimal.Decimal)

	for _, item := range cartItems {
		base := item.Product.BasePrice.Mul(decimal.NewFromInt(int64(item.Quantity)))
		price := base

		// Brand discount
		if strings.ToLower(item.Product.Brand) == "puma" {
			discount := price.Mul(decimal.NewFromFloat(0.4))
			price = price.Sub(discount)
			applied["PUMA 40%"] = applied["PUMA 40%"].Add(discount)
		}

		// Category discount
		if strings.ToLower(item.Product.Category) == "t-shirts" {
			discount := price.Mul(decimal.NewFromFloat(0.1))
			price = price.Sub(discount)
			applied["T-Shirts 10%"] = applied["T-Shirts 10%"].Add(discount)
		}

		originalPrice = originalPrice.Add(base)
		finalPrice = finalPrice.Add(price)
	}

	// Coupon code
	if _, ok := d.validCodes["SUPER69"]; ok {
		discount := finalPrice.Mul(decimal.NewFromFloat(0.69))
		finalPrice = finalPrice.Sub(discount)
		applied["SUPER69"] = discount
	}

	// Bank offer
	if paymentInfo != nil && strings.EqualFold(paymentInfo.BankNameOrDefault(), "ICICI") {
		discount := finalPrice.Mul(decimal.NewFromFloat(0.10))
		finalPrice = finalPrice.Sub(discount)
		applied["ICICI Bank 10%"] = discount
	}

	return &models.DiscountedPrice{
		OriginalPrice:    originalPrice,
		FinalPrice:       finalPrice,
		AppliedDiscounts: applied,
		Message:          "Discounts applied successfully",
	}, nil
}

func (d *discountService) ValidateDiscountCode(ctx context.Context, code string, cartItems []models.CartItem, customer models.CustomerProfile) (bool, error) {
	if _, ok := d.validCodes[strings.ToUpper(code)]; ok {
		return true, nil
	}
	return false, errors.New("invalid or expired discount code")
}

func (p *models.PaymentInfo) BankNameOrDefault() string {
	if p.BankName != nil {
		return *p.BankName
	}
	return ""
}
