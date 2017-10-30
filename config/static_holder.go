package config

type StaticHolder struct {
	V interface{}
}

func (h *StaticHolder) Config() interface{} {
	return h.V
}

func (h *StaticHolder) Close() error {
	return nil
}
