package redisc

type Encoder interface {
	Encode(v interface{}) ([]byte, error)
	Decode(data []byte) (interface{}, error)
}
