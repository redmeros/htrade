package strategies

import "fmt"

// PrintingDataConsumer jest odbiorcą danych
// który drukuje to co dostanie:D - stworzone do testów
type PrintingDataConsumer struct {
}

// OnData tylko wyświetla to co otrzymuje
func (p *PrintingDataConsumer) OnData(data interface{}) {
	fmt.Println(data)
}
