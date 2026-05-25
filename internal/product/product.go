package product

import "time"

type Product struct {
	Name           string
	Volume         float64
	CreatedAt      time.Time
	ExpirationDate int
}
