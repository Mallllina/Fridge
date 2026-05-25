package storage

import (
	"errors"
	"fridge/internal/product"
)

type BaseStorage struct {
	IsOpen        bool
	Products      []product.Product
	MaxVolume     float64
	CurrentVolume float64
}

func (b *BaseStorage) Open() {
	b.IsOpen = true
}

func (b *BaseStorage) Close() {
	b.IsOpen = false
}

func (b *BaseStorage) AddProduct(p product.Product) error {
	if !b.IsOpen {
		return errors.New("Дверь закрыта")
	}

	if p.Volume > b.MaxVolume-b.CurrentVolume {
		return errors.New("Недостаточно места")
	}

	b.Products = append(b.Products, p)
	b.CurrentVolume += p.Volume
	return nil
}

func (b *BaseStorage) RemoveProduct(name string) error {
	index := -1
	for i, prod := range b.Products {
		if prod.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("Продукт не найден")
	}

	remove := b.Products[index]
	b.Products = append(b.Products[:index], b.Products[index+1:]...)
	b.CurrentVolume -= remove.Volume
	return nil
}
