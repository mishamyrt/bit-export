package export

type CipherCard struct {
	CardholderName string `json:"cardholderName,omitempty"`
	Brand          string `json:"brand,omitempty"`
	Number         string `json:"number,omitempty"`
	ExpMonth       string `json:"expMonth,omitempty"`
	ExpYear        string `json:"expYear,omitempty"`
	Code           string `json:"code,omitempty"`
}
