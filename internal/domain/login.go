package domain

type URIMatchType uint8

const (
	DomainMatchType            URIMatchType = 0
	HostMatchType              URIMatchType = 1
	StartsWithMatchType        URIMatchType = 2
	ExactMatchType             URIMatchType = 3
	RegularExpressionMatchType URIMatchType = 4
	NeverMatchType             URIMatchType = 5
)

type URIMatch struct {
	URI  string       `json:"Uri,omitempty"`
	Type URIMatchType `json:"Match,omitempty"`
}

type CipherLogin struct {
	Password string     `json:"Password,omitempty"`
	TokenOTP string     `json:"Totp,omitempty"`
	Username string     `json:"Username,omitempty"`
	URIs     []URIMatch `json:"Uris,omitempty"`
}
