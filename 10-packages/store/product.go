// Packaage store - тестовый пакет
package store

// Структура Product - описывает продукт, который можно купить
// Имя начинается с большой буквы - поэтому Product будет виден
// за пределами пакета store. Поэтому для манипулирования этим полем
// мы в конструкторе предоставим параметр price, а для доступа и изменения
// в дальнейшем - объявим экспортируемые методы Price() и SetPrice()
// На параметры функий и методов правило видимости в зависимости от
// регистра первой буквы имени не распространяется.
type Product struct {
	Name     string
	Quantity int
	// Поле price начинается с маленькой буквы, и поэтому не будет видно
	// за пределами пакета.
	price float64
}

// Если в качестве параметра передаём 0, то величина налога будет браться по
// умолчанию, как определено в tax.go
var standardTax = newTaxRate(0)

func NewProduct(name string, quantity int, price float64) *Product {
	return &Product{name, quantity, price}
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}

func (p *Product) PriceWithTax() float64 {
	return standardTax.calcPriceWithTax(p)
}
