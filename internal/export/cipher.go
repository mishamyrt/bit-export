package export

import "bitexporter/internal/domain"

type Field struct {
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	Type     int    `json:"type,omitempty"`
	LinkedId string `json:"linkedId,omitempty"`
}

// type Field struct {
// 	Name     string `json:"Name,omitempty"`
// 	Value    string `json:"Value,omitempty"`
// 	Type     int    `json:"Type,omitempty"`
// 	LinkedId string `json:"LinkedId,omitempty"`
// }

type Cipher struct {
	Name           string              `json:"name,omitempty"`
	ID             string              `json:"id,omitempty"`
	Fields         []Field             `json:"fields,omitempty"`
	Login          *Login              `json:"login,omitempty"`
	Card           *CipherCard         `json:"card,omitempty"`
	Identity       *CipherIdentity     `json:"identity,omitempty"`
	SecureNote     *CipherSecureNote   `json:"secureNote,omitempty"`
	CollectionIDs  []string            `json:"collectionIds,omitempty"`
	OrganizationId string              `json:"organizationId,omitempty"`
	FolderID       string              `json:"folderId,omitempty"`
	Favorite       bool                `json:"favorite,omitempty"`
	Type           domain.CipherType   `json:"type,omitempty"`
	Notes          string              `json:"notes,omitempty"`
	Reprompt       domain.RepromptType `json:"reprompt,omitempty"`
}
