package models

import (
	"encoding/json"
)

// IStrategy jest interfejsem dla wszystkich strategii
type IStrategy interface {
	CodeName() string
	SetCodeName(val string)

	DisplayName() string
	SetDisplayName(val string)

	Description() string
	SetDescription(val string)

	SetupDefault()
	MarshalJSON() ([]byte, error)

	InitialMoney() float64
	SetInitialMoney(val float64)

	Leverage() float64
	SetLeverage(val float64)

	InputParameters() []InputParameter
	SetInputParameters(vals []InputParameter)

	Run() (*PortfolioResult, error)
}

// StrategyBase implementuje podstawowe/domyślne metody
// dla wszystkich strategii - ale nie musi być obowiązkowo
// używany
type StrategyBase struct {
	initialMoney    float64
	leverage        float64
	codeName        string
	displayName     string
	description     string
	inputParameters []InputParameter
}

// InputParameters to getter
func (s *StrategyBase) InputParameters() []InputParameter {
	return s.inputParameters
}

// SetInputParameters to setter
func (s *StrategyBase) SetInputParameters(val []InputParameter) {
	s.inputParameters = val
}

// InitialMoney zwraca portfel początkowy
func (s *StrategyBase) InitialMoney() float64 {
	return s.initialMoney
}

// SetInitialMoney to setter dla initial money
func (s *StrategyBase) SetInitialMoney(val float64) {
	s.initialMoney = val
}

// Leverage zwraca lewar
func (s *StrategyBase) Leverage() float64 {
	return s.leverage
}

// SetLeverage to setter dla leverage
func (s *StrategyBase) SetLeverage(val float64) {
	s.leverage = val
}

// CodeName zwraca codename
func (s *StrategyBase) CodeName() string {
	return s.codeName
}

// SetCodeName to setter dla codeName
func (s *StrategyBase) SetCodeName(val string) {
	s.codeName = val
}

// DisplayName zwraca nazwe wyswietlana
func (s *StrategyBase) DisplayName() string {
	return s.displayName
}

// SetDisplayName to setter dla displayname
func (s *StrategyBase) SetDisplayName(val string) {
	s.displayName = val
}

// Description zwraca opis
func (s *StrategyBase) Description() string {
	return "description has not been set for this strategy"
}

// SetDescription to setter dla description
func (s *StrategyBase) SetDescription(val string) {
	s.description = val
}

// MarshalJSON konwertuje typ na JSON
func (s *StrategyBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code_name":    s.CodeName(),
		"display_name": s.DisplayName(),
		"description":  s.Description(),
		"parameters":   s.InputParameters(),
	})
}
