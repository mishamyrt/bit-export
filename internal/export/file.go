package export

type Folder struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type File struct {
	Encrypted bool     `json:"encrypted"`
	Folders   []Folder `json:"folders,omitempty"`
	Items     []Cipher `json:"items,omitempty"`
}
