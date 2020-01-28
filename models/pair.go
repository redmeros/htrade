package models

import (
	"encoding/json"
	"fmt"
)

// Pair to model dla pary walutowej `gorm:"unique_index:idx_pair_time"`
type Pair struct {
	ID    uint64 `gorm:"primary_key" json:"id"`
	Major string `gorm:"type:varchar(3);unique_index:idx_major_minor" json:"major" binding:"required"`
	Minor string `gorm:"type:varchar(3);unique_index:idx_major_minor" json:"minor" binding:"required"`
}

// Name zwraca nazwe pary walutowej
func (p *Pair) Name() string {
	return fmt.Sprintf("%s%s", p.Major, p.Minor)
}

// NameWithSep zwraca nazwę pary z separatorem
func (p *Pair) NameWithSep(sep string) string {
	if len(sep) == 0 {
		sep = "_"
	}
	return fmt.Sprintf("%s%s%s", p.Major, sep, p.Minor)
}

// MarshalJSON zwraca json z pary dodajac do niego
// wartosc NAME (readonly)
func (p Pair) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID    uint64 `json:"id"`
		Minor string `json:"minor"`
		Major string `json:"major"`
		Name  string `json:"name"`
	}{
		ID:    p.ID,
		Minor: p.Minor,
		Major: p.Major,
		Name:  p.Name(),
	})
}
