package wallets

import "github.com/redmeros/htrade/internal/strategies"

// TestWallet jest testowym walletem
type TestWallet struct {
}

// FilterAlgosResult to implementacja dla IWallet
func (t *TestWallet) FilterAlgosResult([]*strategies.Position, []*strategies.AlgoResult) ([]*strategies.Order, error) {
	panic("not implemented")
}

// Positions to getter dla pozycji implementacja dla IWallet
func (t *TestWallet) Positions() []*strategies.Position {
	panic("not implemented")
}

// UpdatePositions to implementacja dla IWallet
func (t *TestWallet) UpdatePositions(orders []*strategies.Order) {
	panic("not implemented")
}
