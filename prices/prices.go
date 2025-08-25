package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-cal/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Method to read the prices values.
// This pointer here is imp, so that we can set the value in the InputPrices and not the copy of InputPrices!
func (job *TaxIncludedPriceJob) LoadPricesData() {

	// CONVERSION OF STIRING PRICE TO FLOAT ------
	// prices := make([]float64, len(lines))

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	// CONVERSION OF STIRING PRICE TO FLOAT ------
	job.InputPrices = prices

	file.Close()
}

// Method - Adding a reciever argument makes this method.
func (job *TaxIncludedPriceJob) Process() {

	// LOAD THE DATA FROM FILE ------
	job.LoadPricesData()

	// result := make(map[string]float64) --> Earlier option.

	// After decimal conversion
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.1f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
