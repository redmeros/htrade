package models

import (
	"fmt"
)

// Pair to model dla pary walutowej
type Pair struct {
	ID    uint64 `gorm:"primary_key"`
	Major string `gorm:"type:varchar(3)"`
	Minor string `gorm:"type:varchar(3)"`
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
