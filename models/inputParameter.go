package models

// InputParameter jest typem który co do zasady ma
// być uzupełniany przez użytkownika
type InputParameter interface {
	DisplayName() string
	Name() string
	Type() string
	Value() interface{}
}

// InputParameterBase jest parametrem wejściowym
type InputParameterBase struct {
	PDisplayName string `json:"display_name"`
	PName        string `json:"name"`
	PType        string `json:"type"`
}

// Name zwraca nazwe kodowa
func (p *InputParameterBase) Name() string {
	return p.PName
}

// Type zwraca typ przechowywanej wartości
func (p *InputParameterBase) Type() string {
	return p.PType
}

// DisplayName zwraca nazwę wyświetlaną użytkownikowi
func (p *InputParameterBase) DisplayName() string {
	return p.PDisplayName
}

// FloatParameter jest parametrem wejściowym typu float
type FloatParameter struct {
	InputParameterBase
	PValue float64 `json:"value"`
}

// Value zwraca wartosc float
func (p *FloatParameter) Value() interface{} {
	return p.PValue
}

// NewFloatParameter zwraca nowy input parameter typu float
func NewFloatParameter(displayName string, name string, value float64) InputParameter {
	w := FloatParameter{}
	w.PDisplayName = displayName
	w.PName = name
	w.PType = "float"
	w.PValue = value
	return &w
}

type dstruct struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// DateParameter jest parametrem wejsciowym typu struct{}
// zawierajacym date
type DateParameter struct {
	InputParameterBase
	PValue dstruct `json:"value"`
}

// Value zwraca wartość typu dstruct
func (p *DateParameter) Value() interface{} {
	return p.PValue
}

// NewDateParameter zwraca nowy input parameter typu dstruct
func NewDateParameter(displayName string, name string, year, month, day int) InputParameter {
	w := DateParameter{}
	w.PDisplayName = displayName
	w.PName = name
	w.PType = "date"
	w.PValue = dstruct{Year: year, Month: month, Day: day}
	return &w
}
