package main

import "fmt"

type DiscountStrategy func(float64) float64

func CalculateDiscount(amount float64, strategy DiscountStrategy) float64 {
	return strategy(amount)
}

func NoDiscount(amount float64) float64 {
	return amount
}

func FixedDiscount(discountAmount float64) DiscountStrategy {
	return func(amount float64) float64 {
		return amount - discountAmount
	}
}

func main() {
	amountWithoutDiscount := 100.0
	fmt.Printf("Сумма без скидки: $%.2f\n", amountWithoutDiscount)

	amountWithFixedDiscount := CalculateDiscount(amountWithoutDiscount, FixedDiscount(10.0))
	fmt.Printf("Сумма со скидкой: $%.2f\n", amountWithFixedDiscount)
}
