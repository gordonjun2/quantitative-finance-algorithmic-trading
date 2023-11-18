package main

import (
	"fmt"
	"math"
)

type CouponBond struct {
	principal    float64
	rate         float64
	maturity     float64
	interestRate float64
}

func NewCouponBond(principal, rate, maturity, interestRate float64) *CouponBond {
	return &CouponBond{
		principal:    principal,
		rate:         rate / 100,
		maturity:     maturity,
		interestRate: interestRate / 100,
	}
}

func (bond *CouponBond) presentValue(x float64, n float64) float64 {
	return x / math.Pow(1+bond.interestRate, n)
}

func (bond *CouponBond) calculatePrice() float64 {

	price := 0.0
	couponAmount := bond.principal * bond.rate

	for i := 1.0; i <= bond.maturity; i++ {
		price += bond.presentValue(couponAmount, i)
	}

	price += bond.presentValue(bond.principal, bond.maturity)

	return price
}

func main() {
	principal := 1000.0
	rate := 10.0
	maturity := 3.0
	interestRate := 4.0

	bond := NewCouponBond(principal, rate, maturity, interestRate)
	price := bond.calculatePrice()

	fmt.Printf("Price of the coupon bond: %.2f\n", price)
}
