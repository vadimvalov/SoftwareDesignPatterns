package main

import (
	"fmt"
)

type Shawarma interface {
	Description() string
}

type ChickenShawarma struct{}

func (s *ChickenShawarma) Description() string {
	return "Куриная шаурма"
}

type BeefShawarma struct{}

func (s *BeefShawarma) Description() string {
	return "Говяжья шаурма (пахнет собакой)"
}

type ShawarmaFactory interface {
	CreateShawarma() Shawarma
}

type ChickenShawarmaFactory struct{}

func (f *ChickenShawarmaFactory) CreateShawarma() Shawarma {
	return &ChickenShawarma{}
}

type BeefShawarmaFactory struct{}

func (f *BeefShawarmaFactory) CreateShawarma() Shawarma {
	return &BeefShawarma{}
}

func main() {
	chickenShawarmaFactory := &ChickenShawarmaFactory{}
	chickenShawarma := chickenShawarmaFactory.CreateShawarma()
	fmt.Println("Продукт из первой лавки:", chickenShawarma.Description())

	beefShawarmaFactory := &BeefShawarmaFactory{}
	beefShawarma := beefShawarmaFactory.CreateShawarma()
	fmt.Println("Продукт из лавки у вокзала:", beefShawarma.Description())
}
