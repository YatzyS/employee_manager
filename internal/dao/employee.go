package dao

type Employee struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}
