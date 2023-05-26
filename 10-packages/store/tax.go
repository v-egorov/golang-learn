package store

// никакие фичи из tax.go не экспортируются за пределы пакета store
const defaultTaxRate float64 = 0.2

type taxRate struct {
	rate float64
}

func newTaxRate(rate float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	return &taxRate{rate}
}

func (taxRate *taxRate) calcPriceWithTax(product *Product) float64 {
	// можем использовать не экспортированную фичу product.price - потому
	// что мы в том же пакете
	// Кроме этого, не нужно ссылаться на ptoduct с префиксом
	// названия пакета - как store.product
	return product.price + product.price*taxRate.rate
}
