package domain

type Folder struct {
	Id   string
	Name string
}

type Sync struct {
	Ciphers []Cipher
	// Collections []interface{}
	// Domains     interface{}
	Folders []Folder
	Profile Profile
	// Sends       interface{}
}
