package domain

type CipherIdentity struct {
	Title          string `json:"Title,omitempty"`
	FirstName      string `json:"FirstName,omitempty"`
	MiddleName     string `json:"MiddleName,omitempty"`
	LastName       string `json:"LastName,omitempty"`
	Address1       string `json:"Address1,omitempty"`
	Address2       string `json:"Address2,omitempty"`
	Address3       string `json:"Address3,omitempty"`
	City           string `json:"City,omitempty"`
	State          string `json:"State,omitempty"`
	PostalCode     string `json:"PostalCode,omitempty"`
	Country        string `json:"Country,omitempty"`
	Company        string `json:"Company,omitempty"`
	Email          string `json:"Email,omitempty"`
	Phone          string `json:"Phone,omitempty"`
	SSN            string `json:"SSN,omitempty"`
	Username       string `json:"Username,omitempty"`
	PassportNumber string `json:"PassportNumber,omitempty"`
	LicenseNumber  string `json:"LicenseNumber,omitempty"`
}
