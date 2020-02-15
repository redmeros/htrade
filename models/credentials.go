package models

// Credentials zawiera dane logowania
//
type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetIdentifier zwraca email lub haslo
// w zaleznosci co jest ustawione
func (c *Credentials) GetIdentifier() string {
	if c.Email != "" {
		return c.Email
	} else if c.Username != "" {
		return c.Username
	}
	return ""
}
