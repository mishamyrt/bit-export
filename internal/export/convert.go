package export

import (
	"bit-exporter/internal/domain"
)

func FromDomain(sync *domain.Sync) File {
	var file File
	file.Folders = make([]Folder, len(sync.Folders))
	for i := range sync.Folders {
		file.Folders[i] = Folder(sync.Folders[i])
	}
	for _, c := range sync.Ciphers {
		if c.DeletedDate != nil {
			continue
		}
		cipher := Cipher{
			Name:          c.Name,
			ID:            c.ID,
			Fields:        make([]Field, len(c.Fields)),
			FolderID:      c.FolderID,
			CollectionIDs: c.CollectionIDs,
			Favorite:      c.Favorite,
			Type:          c.Type,
			Notes:         c.Notes,
			Reprompt:      c.Reprompt,
		}
		for j := range c.Fields {
			cipher.Fields[j] = Field(c.Fields[j])
		}

		switch c.Type {
		case domain.LoginCipherType:
			cipher.Login = &Login{
				Username: c.Login.Username,
				Password: c.Login.Password,
				TokenOTP: c.Login.TokenOTP,
				URIs:     make([]URIMatch, len(c.Login.URIs)),
			}
			for j := range c.Login.URIs {
				cipher.Login.URIs[j] = URIMatch(c.Login.URIs[j])
			}
		case domain.CardCipherType:
			cipher.Card = (*CipherCard)(c.Card)
		case domain.IdentityCipherType:
			cipher.Identity = (*CipherIdentity)(c.Identity)
		case domain.SecureNoteCipherType:
			cipher.SecureNote = (*CipherSecureNote)(c.SecureNote)
		}
		file.Items = append(file.Items, cipher)
	}
	return file
}
