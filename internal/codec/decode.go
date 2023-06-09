package codec

import (
	"bit-exporter/internal/domain"
	"bit-exporter/internal/export"
	"bit-exporter/pkg/crypto"
)

func (c *Codec) DecodeString(target *string) error {
	cipher := crypto.CipherFromString(*target)
	message, err := crypto.Decrypt(cipher, c.key, c.keyMac)
	if err != nil {
		return err
	}
	*target = string(message)
	return nil
}

func (c *Codec) tryDecode(pointers ...*string) error {
	var err error
	for i := range pointers {
		if len(*pointers[i]) > 0 {
			err = c.DecodeString(pointers[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Codec) Decode(file *export.File) error {
	err := c.decodeFolders(file.Folders)
	if err != nil {
		return err
	}
	err = c.decodeCiphers(file.Items)
	if err != nil {
		return err
	}
	return nil
}

func (c *Codec) decodeFolders(folders []export.Folder) error {
	for i := range folders {
		err := c.DecodeString(&folders[i].Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Codec) decodeCiphers(ciphers []export.Cipher) error {
	for i := range ciphers {
		err := c.tryDecode(
			&ciphers[i].Name,
			&ciphers[i].Notes,
		)
		if err != nil {
			return err
		}
		err = c.decodeFields(ciphers[i].Fields)
		if err != nil {
			return err
		}

		switch ciphers[i].Type {
		case domain.LoginCipherType:
			c.decodeLogin(ciphers[i].Login)
		case domain.CardCipherType:
			c.decodeCard(ciphers[i].Card)
		case domain.IdentityCipherType:
			c.decodeIdentity(ciphers[i].Identity)
		}
	}
	return nil
}

func (c *Codec) decodeFields(fields []export.Field) error {
	for i := range fields {
		err := c.tryDecode(&fields[i].Name, &fields[i].Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Codec) decodeCard(note *export.CipherCard) error {
	return c.tryDecode(
		&note.Brand,
		&note.CardholderName,
		&note.Code,
		&note.ExpMonth,
		&note.ExpYear,
		&note.Number,
	)
}

func (c *Codec) decodeIdentity(id *export.CipherIdentity) error {
	return c.tryDecode(
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
	)
}

func (c *Codec) decodeLogin(login *export.Login) error {
	err := c.tryDecode(
		&login.Username,
		&login.Password,
		&login.TokenOTP,
	)
	if err != nil {
		return err
	}
	for i := range login.URIs {
		err := c.DecodeString(&login.URIs[i].URI)
		if err != nil {
			return err
		}
	}
	return nil
}
