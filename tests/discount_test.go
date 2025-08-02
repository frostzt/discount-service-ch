package tests

import (
	"context"
	"strings"
	"testing"

	"github.com/frostzt/discount/models"
	"github.com/frostzt/discount/service"
	"github.com/frostzt/discount/testdata"
)

func TestCalculateCartDiscounts(t *testing.T) {
	svc := service.NewDiscountService()
	cart := testdata.FakeCartData()
	customer := testdata.FakeCustomer()
	payment := testdata.FakePaymentInfo()

	discounted, err := svc.CalculateCartDiscounts(context.Background(), cart, customer, payment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if discounted.OriginalPrice.LessThanOrEqual(discounted.FinalPrice) {
		t.Errorf("Final price should be less than original after discounts")
	}

	if len(discounted.AppliedDiscounts) == 0 {
		t.Error("No discounts were applied")
	}

	t.Logf("Original: %v | Final: %v", discounted.OriginalPrice, discounted.FinalPrice)
	for k, v := range discounted.AppliedDiscounts {
		t.Logf("%s -> %v", k, v)
	}
}

func TestValidateDiscountCode(t *testing.T) {
	svc := service.NewDiscountService()

	tests := []struct {
		name      string
		code      string
		cart      []models.CartItem
		customer  models.CustomerProfile
		expectOK  bool
		expectErr string
	}{
		{
			name:     "Valid SUPER69 with gold user on PUMA T-shirt",
			code:     "SUPER69",
			cart:     testdata.FakeCartData(),
			customer: testdata.FakeCustomer(),
			expectOK: true,
		},
		{
			name:      "Invalid SUPER69 with Nike brand",
			code:      "SUPER69",
			cart:      testdata.FakeNikeCartItem(),
			customer:  testdata.FakeCustomer(),
			expectOK:  false,
			expectErr: "brand",
		},
		{
			name:      "Invalid SUPER69 with silver user",
			code:      "SUPER69",
			cart:      testdata.FakeCartData(),
			customer:  testdata.FakeSilverCustomer(),
			expectOK:  false,
			expectErr: "tier",
		},
		{
			name:      "Invalid fake code",
			code:      "FAKECODE",
			cart:      testdata.FakeCartData(),
			customer:  testdata.FakeCustomer(),
			expectOK:  false,
			expectErr: "invalid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := svc.ValidateDiscountCode(context.Background(), tc.code, tc.cart, tc.customer)
			if ok != tc.expectOK {
				t.Errorf("expected ok=%v, got %v", tc.expectOK, ok)
			}
			if err != nil && tc.expectErr != "" && !strings.Contains(err.Error(), tc.expectErr) {
				t.Errorf("expected error to contain '%s', got '%v'", tc.expectErr, err)
			}
		})
	}
}
