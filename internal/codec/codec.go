package codec

type Codec struct {
	key    []byte
	keyMac []byte
}

func (c *Codec) SetKeys(key []byte, keyMac []byte) {
	c.key = key
	c.keyMac = keyMac
}

func New(key []byte, keyMac []byte) *Codec {
	return &Codec{
		key:    key,
		keyMac: keyMac,
	}
}
