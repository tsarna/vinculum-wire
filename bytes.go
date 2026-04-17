package wire

type bytesFormat struct{}

func (bytesFormat) Serialize(v any) ([]byte, error) {
	s, err := scalarToString(v)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (bytesFormat) SerializeString(v any) (string, error) {
	return scalarToString(v)
}

func (bytesFormat) Deserialize(b []byte) (any, error) {
	out := make([]byte, len(b))
	copy(out, b)
	return out, nil
}

func (bytesFormat) Name() string { return "bytes" }
