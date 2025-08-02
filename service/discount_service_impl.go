package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/frostzt/discount/models"
	"github.com/shopspring/decimal"
)

type brandDiscount struct {
	Name    string
	Brand   string
	Percent float64
}

type categoryDiscount struct {
	Name     string
	Category string
	Percent  float64
}

type bankOffer struct {
	Name    string
	Bank    string
	Percent float64
}

type couponDiscount struct {
	Code              string
	Name              string
	Percent           float64
	AllowedBrands     []string
	ExcludedBrands    []string
	AllowedCategories []string
	MinCustomerTier   string
}

type discountService struct {
	brandDiscounts    []brandDiscount
	categoryDiscounts []categoryDiscount
	couponDiscounts   []couponDiscount
	bankOffers        []bankOffer
}

func NewDiscountService() DiscountService {
	return &discountService{
		brandDiscounts: []brandDiscount{
			{Name: "PUMA 40%", Brand: "puma", Percent: 0.40},
			{Name: "Nike 30%", Brand: "nike", Percent: 0.30},
		},
		categoryDiscounts: []categoryDiscount{
			{Name: "T-Shirts 10%", Category: "t-shirts", Percent: 0.10},
			{Name: "Jeans 20%", Category: "jeans", Percent: 0.20},
		},
		couponDiscounts: []couponDiscount{
			{
				Code:              "SUPER69",
				Name:              "SUPER69",
				Percent:           0.69,
				ExcludedBrands:    []string{"Nike"},
				AllowedCategories: []string{"t-shirts"},
				MinCustomerTier:   "gold",
			},
			{
				Code:              "SUMMER50",
				Name:              "SUMMER50",
				Percent:           0.50,
				AllowedCategories: []string{"shorts", "t-shirts"},
				MinCustomerTier:   "silver",
			},
		},
		bankOffers: []bankOffer{
			{Name: "ICICI Bank 10%", Bank: "ICICI", Percent: 0.10},
			{Name: "HDFC Debit 5%", Bank: "HDFC", Percent: 0.05},
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
		for _, bd := range d.brandDiscounts {
			if strings.EqualFold(item.Product.Brand, bd.Brand) {
				discount := price.Mul(decimal.NewFromFloat(bd.Percent))
				price = price.Sub(discount)
				applied[bd.Name] = applied[bd.Name].Add(discount)
			}
		}

		// Category discount
		for _, cd := range d.categoryDiscounts {
			if strings.EqualFold(item.Product.Category, cd.Category) {
				discount := price.Mul(decimal.NewFromFloat(cd.Percent))
				price = price.Sub(discount)
				applied[cd.Name] = applied[cd.Name].Add(discount)
			}
		}

		originalPrice = originalPrice.Add(base)
		finalPrice = finalPrice.Add(price)
	}

	// Coupon code
	for _, coupon := range d.couponDiscounts {
		// Optional: check if code is allowed based on cart, etc.
		discount := finalPrice.Mul(decimal.NewFromFloat(coupon.Percent))
		finalPrice = finalPrice.Sub(discount)
		applied[coupon.Name] = discount
	}

	// Bank offer
	if paymentInfo != nil {
		for _, offer := range d.bankOffers {
			if strings.EqualFold(paymentInfo.BankNameOrDefault(), offer.Bank) {
				discount := finalPrice.Mul(decimal.NewFromFloat(offer.Percent))
				finalPrice = finalPrice.Sub(discount)
				applied[offer.Name] = discount
			}
		}
	}
	return &models.DiscountedPrice{
		OriginalPrice:    originalPrice,
		FinalPrice:       finalPrice,
		AppliedDiscounts: applied,
		Message:          "Discounts applied successfully",
	}, nil
}

func (d *discountService) ValidateDiscountCode(
	ctx context.Context,
	code string,
	cartItems []models.CartItem,
	customer models.CustomerProfile,
) (bool, error) {
	var matched *couponDiscount
	for _, coupon := range d.couponDiscounts {
		if strings.EqualFold(coupon.Code, code) {
			matched = &coupon
			break
		}
	}
	if matched == nil {
		return false, errors.New("invalid discount code")
	}

	for _, item := range cartItems {
		brand := strings.ToLower(item.Product.Brand)
		category := strings.ToLower(item.Product.Category)

		// Excluded brand
		for _, excl := range matched.ExcludedBrands {
			if brand == strings.ToLower(excl) {
				return false, fmt.Errorf("discount not valid for brand %s", item.Product.Brand)
			}
		}

		// Allowed category
		if len(matched.AllowedCategories) > 0 {
			ok := false
			for _, allowed := range matched.AllowedCategories {
				if category == strings.ToLower(allowed) {
					ok = true
					break
				}
			}
			if !ok {
				return false, fmt.Errorf("discount not valid for category %s", item.Product.Category)
			}
		}
	}

	// Optional: check customer tier
	if matched.MinCustomerTier != "" && !strings.EqualFold(customer.Tier, matched.MinCustomerTier) {
		return false, fmt.Errorf("discount requires customer tier %s", matched.MinCustomerTier)
	}

	return true, nil
}
