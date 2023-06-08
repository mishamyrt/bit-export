package crypto

import (
	"bit-exporter/internal/domain"
)

type Coder struct {
	key    []byte
	macKey []byte
}

func (c *Coder) SetKeys(key []byte, macKey []byte) {
	c.key = key
	c.macKey = macKey
}

func (c *Coder) DecryptString(payload *string) error {
	cipher := FromString(*payload)
	message, err := decryptWith(cipher, c.key, c.macKey)
	if err != nil {
		return err
	}
	*payload = string(message)
	return nil
}

func (c *Coder) decryptStrings(pointers []*string) error {
	var err error
	for i := range pointers {
		if len(*pointers[i]) > 0 {
			err = c.DecryptString(pointers[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Coder) Decrypt(payload []byte) ([]byte, error) {
	cipher := FromBytes(payload)
	return decryptWith(cipher, d.key, d.macKey)
}

func (d *Coder) DecryptSync(sync *domain.Sync) error {
	err := d.decryptFolders(sync.Folders)
	if err != nil {
		return err
	}
	err = d.decryptCiphers(sync.Ciphers)
	if err != nil {
		return err
	}
	return nil
}

func (d *Coder) decryptFolders(folders []domain.Folder) error {
	for i := range folders {
		err := d.DecryptString(&folders[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Coder) decryptCiphers(ciphers []domain.Cipher) error {
	for i := range ciphers {
		err := c.DecryptString(&ciphers[i].Name)
		if err != nil {
			return err
		}
		err = c.decryptFields(ciphers[i].Fields)
		if err != nil {
			return err
		}

		switch ciphers[i].Type {
		case domain.LoginCipherType:
			c.decryptLogin(ciphers[i].Login)
		case domain.CardCipherType:
			c.decryptCard(ciphers[i].Card)
		case domain.IdentityCipherType:
			c.decryptIdentity(ciphers[i].Identity)
		}
		// 	LoginCipherType      CipherType = 1
		// SecureNoteCipherType CipherType = 2
		// CardCipherType       CipherType = 3
		// IdentityCipherType   CipherType = 4
	}
	return nil
}

func (c *Coder) decryptFields(fields []domain.Field) error {
	for i := range fields {
		err := c.decryptStrings([]*string{
			&fields[i].Name,
			&fields[i].Value,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Coder) decryptCard(note *domain.CipherCard) error {
	return c.decryptStrings([]*string{
		&note.Brand,
		&note.CardholderName,
		&note.Code,
		&note.ExpMonth,
		&note.ExpYear,
		&note.Number,
	})
}

func (c *Coder) decryptIdentity(id *domain.CipherIdentity) error {
	return c.decryptStrings([]*string{
		&id.Address1,
		&id.Address2,
		&id.Address3,
		&id.City,
		&id.Company,
		&id.Country,
		&id.Email,
		&id.FirstName,
		&id.LastName,
		&id.MiddleName,
		&id.PassportNumber,
		&id.Phone,
		&id.PostalCode,
		&id.SSN,
		&id.State,
		&id.Title,
		&id.Username,
	})
}

func (c *Coder) decryptLogin(login *domain.CipherLogin) error {
	err := c.decryptStrings([]*string{
		&login.Username,
		&login.Password,
		&login.TokenOTP,
	})
	if err != nil {
		return err
	}
	for i := range login.URIs {
		err := c.DecryptString(&login.URIs[i].URI)
		if err != nil {
			return err
		}
	}
	return nil
}
