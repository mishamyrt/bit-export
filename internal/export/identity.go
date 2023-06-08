package export

type CipherIdentity struct {
	Title          string `json:"title,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	MiddleName     string `json:"middleName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	Address1       string `json:"address1,omitempty"`
	Address2       string `json:"address2,omitempty"`
	Address3       string `json:"address3,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	PostalCode     string `json:"postalCode,omitempty"`
	Country        string `json:"country,omitempty"`
	Company        string `json:"company,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	SSN            string `json:"ssn,omitempty"`
	Username       string `json:"username,omitempty"`
	PassportNumber string `json:"passportNumber,omitempty"`
	LicenseNumber  string `json:"licenseNumber,omitempty"`
}
