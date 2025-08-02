package service

import (
	"context"

	"github.com/frostzt/discount/models"
)

type DiscountService interface {
	CalculateCartDiscounts(
		ctx context.Context,
		cartItems []models.CartItem,
		customer models.CustomerProfile,
		paymentInfo *models.PaymentInfo) (*models.DiscountedPrice, error)

	ValidateDiscountCode(
		ctx context.Context,
		code string,
		cartItems []models.CartItem,
		customer models.CustomerProfile) (bool, error)
}
