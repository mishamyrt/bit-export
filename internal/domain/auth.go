package domain

type Auth struct {
	Kdf            uint8
	KdfIterations  int
	KdfMemory      int
	KdfParallelism int
	Key            string
	PrivateKey     string
	AccessToken    string `json:"access_token"`
}
