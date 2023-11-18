package main

import (
	"fmt"
	"math"
)

type ZeroCouponBond struct {
	principal    float64
	maturity     float64
	interestRate float64
}

func NewZeroCouponBond(principal, maturity, interestRate float64) *ZeroCouponBond {
	return &ZeroCouponBond{
		principal:    principal,
		maturity:     maturity,
		interestRate: interestRate / 100,
	}
}

func (bond *ZeroCouponBond) presentValue(x float64, n float64) float64 {
	return x / math.Pow(1+bond.interestRate, n)
}

func (bond *ZeroCouponBond) calculatePrice() float64 {
	return bond.presentValue(bond.principal, bond.maturity)
}

func main() {
	principal := 1000.0
	maturity := 2.0
	interestRate := 4.0

	bond := NewZeroCouponBond(principal, maturity, interestRate)
	price := bond.calculatePrice()

	fmt.Printf("Price of the zero-coupon bond: %.2f\n", price)
}
