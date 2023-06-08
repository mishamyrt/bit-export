package export

import (
	"bit-exporter/internal/domain"
)

type URIMatch struct {
	URI  string              `json:"uri,omitempty"`
	Type domain.URIMatchType `json:"match,omitempty"`
}

type Login struct {
	URIs     []URIMatch `json:"uris"`
	Username string     `json:"username,omitempty"`
	Password string     `json:"password,omitempty"`
	TokenOTP string     `json:"totp,omitempty"`
}
