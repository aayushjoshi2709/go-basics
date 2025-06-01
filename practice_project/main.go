package main

import (
	"aayushjoshi2709/practice_project/filemanager"
	"aayushjoshi2709/practice_project/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for taxRateIndex, taxRate := range taxRates {
		doneChans[taxRateIndex] = make(chan bool)
		errorChans[taxRateIndex] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[taxRateIndex], errorChans[taxRateIndex])

		// cm := cmdmanager.New()
		// priceJob = prices.NewTaxIncludedPriceJob(cm, taxRate)
		// go priceJob.Process(doneChans[taxRateIndex], errorChans[taxRateIndex])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!!!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
