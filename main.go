package main

import (
	"example.com/price-cal/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// var result map[float64][]float64 = make(map[float64][]float64)
	// Short-hand way of writing the above line -
	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
