package main 

type CurrencyConversionService struct {
	rates map[string]map[string]float64
}


func (ccs *CurrencyConversionService) Convert(from Currency, to Currency, amount float64) float64 {
	if from.code == to.code {
		return amount
	} 
	return amount * ccs.rates[from.code][to.code]
} 