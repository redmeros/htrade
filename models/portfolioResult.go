package models

// PortfolioResult zawiera informacje
// o portfelu w czasie testów
type PortfolioResult struct {
	Records   []*PortfolioRecord
	Positions []*Position
	MaxValue  float64
	MinValue  float64

	started bool
}

// AddRecord dodaje do PortfolioResult kolejny record,
// I oblicza niezbędne dane
func (p *PortfolioResult) AddRecord(r *PortfolioRecord) {
	p.Records = append(p.Records, r)
	if !p.started {
		p.started = true
		p.MinValue = r.Value
		p.MaxValue = r.Value
	} else {
		if r.Value < p.MinValue {
			p.MinValue = r.Value
		}
		if r.Value > p.MaxValue {
			p.MaxValue = r.Value
		}
	}
}

// AddPosition dodaje do PortfoilioResult kolejną
// Pozycję
func (p *PortfolioResult) AddPosition(r *Position) {
	p.Positions = append(p.Positions, r)
}
