package domain

type CipherCard struct {
	CardholderName string `json:"CardholderName,omitempty"`
	Brand          string `json:"Brand,omitempty"`
	Number         string `json:"Number,omitempty"`
	ExpMonth       string `json:"ExpMonth,omitempty"`
	ExpYear        string `json:"ExpYear,omitempty"`
	Code           string `json:"Code,omitempty"`
}
