package main

import (
	"fmt"

	"github.com/yubinex/go-price-calculator/filemanager"
	"github.com/yubinex/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for idx, taxRate := range taxRates {
		doneChans[idx] = make(chan bool)
		errorChans[idx] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[idx], errorChans[idx])

	}

	for idx := range taxRates {
		select {
		case err := <-errorChans[idx]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[idx]:
			fmt.Println("Done!")
		}
	}
}
