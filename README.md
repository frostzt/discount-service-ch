# Discount Service

This is a Go-based discount engine designed to simulate e-commerce style cart discount scenarios. The service is capable of handling brand-specific discounts, category-based offers, coupon codes, and bank payment offers.

## 📦 Features

- Apply multiple brand-level and category-level discounts
- Support for voucher/coupon codes with restrictions
- Dynamic bank offers
- Cleanly separated and extensible discount logic
- Comprehensive unit tests with test data

## 🧠 Technical Decisions

### 1. **Dynamic Discount Rules**

All discounts are configured via Go structs rather than hardcoded `if` statements. This allows us to:
- Add multiple brand/category/coupon/bank discounts
- Change values without touching business logic

Each rule type is modeled via its own struct:
- `brandDiscount`
- `categoryDiscount`
- `couponDiscount` (with rules)
- `bankOffer`

NOTE: This should generally be in a database

### 2. **Coupon Validation Logic**

`ValidateDiscountCode()` uses metadata to enforce:
- Brand exclusions (e.g., not valid on Nike)
- Category restrictions (e.g., valid only for "T-shirts")
- Minimum customer tier (e.g., only for "gold")

### 3. **Clean API Contract**

Service interface defined as:
```go
CalculateCartDiscounts(ctx, cartItems, customer, paymentInfo) (*DiscountedPrice, error)
ValidateDiscountCode(ctx, code, cartItems, customer) (bool, error)
```

### 4. **Test Data Isolation**

`testdata/` contains reusable fake data:
- Gold & Silver customers
- PUMA & Nike items
- PaymentInfo for bank offers

---

## ✅ How to Run

```bash
go mod tidy
go test ./...
```

## 🧪 Test Coverage

Includes tests for:
- Valid/invalid coupon codes
- Excluded brands
- Tier mismatches
- Cart price discounting

---

## 🗂️ File Structure

```
├── models/              # Core data models
├── service/             # Discount logic
├── testdata/            # Fake cart/customers/payment data
├── tests/               # Unit tests
├── go.mod               # Dependencies
├── README.md            # This file
```
