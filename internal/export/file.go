package export

type Folder struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type File struct {
	Key       *string  `json:"encKeyValidation_DO_NOT_EDIT,omitempty"`
	Encrypted bool     `json:"encrypted"`
	Folders   []Folder `json:"folders,omitempty"`
	Items     []Cipher `json:"items,omitempty"`
}
