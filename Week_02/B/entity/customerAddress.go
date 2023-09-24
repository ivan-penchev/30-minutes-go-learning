package entity

type CustomerAddress struct {
	Name    string `json:"name"` // key
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	Line3   string `json:"line3"`
	Line4   string `json:"line4"`
	Code    string `json:"code"`
	Country string `json:"country"`
}
