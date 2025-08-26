package main

import (
	"fmt"

	"example.com/price-cal/filemanager"
	"example.com/price-cal/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// var result map[float64][]float64 = make(map[float64][]float64)
	// Short-hand way of writing the above line -
	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
