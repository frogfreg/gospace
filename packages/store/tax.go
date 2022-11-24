package store

const defaultTaxRate = 0.2
const minThreshold = 10

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

func init() {
	for category, price := range categoryMaxPrices {

		categoryMaxPrices[category] = price + (price * defaultTaxRate)

	}

}

type taxRate struct {
	rate, threshold float64
}

func NewTaxRate(rate, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}

	if threshold < minThreshold {
		threshold = minThreshold

	}

	return &taxRate{rate, threshold}

}

func (t *taxRate) calcTax(prod *Product) float64 {

	var price float64

	if prod.price > t.threshold {

		price = prod.price + (prod.price * t.rate)

	} else {

		price = prod.price

	}

	if max, ok := categoryMaxPrices[prod.Category]; ok && price > max {

		price = max

	}

	return price

}
