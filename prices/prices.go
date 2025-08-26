package prices

import (
	"fmt"

	"example.com/price-cal/conversion"
	"example.com/price-cal/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

// Method to read the prices values.
// This pointer here is imp, so that we can set the value in the InputPrices and not the copy of InputPrices!
func (job *TaxIncludedPriceJob) LoadPricesData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	// CONVERSION OF STIRING PRICE TO FLOAT ------
	// prices := make([]float64, len(lines))

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	// CONVERSION OF STIRING PRICE TO FLOAT ------
	job.InputPrices = prices
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
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

	job.TaxIncludedPrices = result

	// filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
