package main

import (
	"fmt"
)

type Shawarma interface {
	Price() int
}

type SimpleShawarma struct{}

func (c *SimpleShawarma) Price() int {
	return 5
}

type ShawarmaDecorator struct {
	shawarma Shawarma
}

func (cd *ShawarmaDecorator) Price() int {
	return cd.shawarma.Price()
}

type CheeseDecorator struct {
	shawarma Shawarma
}

func (sd *CheeseDecorator) Price() int {
	return sd.shawarma.Price() + 2
}

type JalapenjoDecorator struct {
	shawarma Shawarma
}

func (md *JalapenjoDecorator) Price() int {
	return md.shawarma.Price() + 3
}

func main() {
	shawarma := &SimpleShawarma{}

	shawarmaWithCheese := &CheeseDecorator{shawarma}
	shawarmaWithJalapenjo := &JalapenjoDecorator{shawarma}

	fmt.Println("Цена обычной шавы:", shawarma.Price())
	fmt.Println("Цена шавы с сыром:", shawarmaWithCheese.Price())
	fmt.Println("Цена шавы с халапеньо:", shawarmaWithJalapenjo.Price())
}
