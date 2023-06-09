package domain

type Field struct {
	Name     string `json:"Name,omitempty"`
	Value    string `json:"Value,omitempty"`
	Type     int    `json:"Type,omitempty"`
	LinkedId string `json:"LinkedId,omitempty"`
}

type CipherType uint8

const (
	LoginCipherType      CipherType = 1
	SecureNoteCipherType CipherType = 2
	CardCipherType       CipherType = 3
	IdentityCipherType   CipherType = 4
)

type RepromptType uint8

const (
	NoneRepromptType     RepromptType = 0
	PasswordRepromptType RepromptType = 1
)

type Cipher struct {
	Name           string            `json:"Name,omitempty"`
	ID             string            `json:"Id,omitempty"`
	Fields         []Field           `json:"Fields,omitempty"`
	Login          *CipherLogin      `json:"Login,omitempty"`
	Card           *CipherCard       `json:"Card,omitempty"`
	Identity       *CipherIdentity   `json:"Identity,omitempty"`
	SecureNote     *CipherSecureNote `json:"SecureNote,omitempty"`
	CollectionIDs  []string          `json:"CollectionIds,omitempty"`
	OrganizationId string            `json:"OrganizationId,omitempty"`
	FolderID       string            `json:"FolderId,omitempty"`
	Favorite       bool              `json:"Favorite"`
	Type           CipherType        `json:"Type,omitempty"`
	Notes          string            `json:"Notes,omitempty"`
	DeletedDate    *string           `json:"DeletedDate,omitempty"`
	// OrganizationUseTotp bool `json:"OrganizationUseTotp,omitempty"`

	Reprompt RepromptType `json:"Reprompt,omitempty"`
}
