package export

import (
	"bit-exporter/internal/domain"
)

func FromDomain(sync *domain.Sync) File {
	var file File
	file.Folders = make([]Folder, len(sync.Folders))
	file.Items = make([]Cipher, len(sync.Ciphers))
	for i := range sync.Folders {
		file.Folders[i] = Folder(sync.Folders[i])
	}
	for i, c := range sync.Ciphers {
		file.Items[i] = Cipher{
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
			file.Items[i].Fields[j] = Field(c.Fields[j])
		}

		switch c.Type {
		case domain.LoginCipherType:
			file.Items[i].Login = &Login{
				Username: c.Login.Username,
				Password: c.Login.Password,
				TokenOTP: c.Login.TokenOTP,
				URIs:     make([]URIMatch, len(c.Login.URIs)),
			}
			for j := range c.Login.URIs {
				file.Items[i].Login.URIs[j] = URIMatch(c.Login.URIs[j])
			}
		case domain.CardCipherType:
			file.Items[i].Card = (*CipherCard)(c.Card)
		case domain.IdentityCipherType:
			file.Items[i].Identity = (*CipherIdentity)(c.Identity)
		case domain.SecureNoteCipherType:
			file.Items[i].SecureNote = (*CipherSecureNote)(c.SecureNote)
		}
		// file.Items[i].Login = Login(c.Login)
	}
	return file
}
