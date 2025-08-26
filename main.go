package main

import (
	"example.com/price-cal/cmdmanager"
	"example.com/price-cal/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// var result map[float64][]float64 = make(map[float64][]float64)
	// Short-hand way of writing the above line -
	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		cmdm := cmdmanager.New()

		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
