package wire

import (
	"encoding/json"
	"fmt"
)

type jsonFormat struct{}

func (jsonFormat) Serialize(v any) ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}
	if b, ok := v.([]byte); ok {
		return b, nil
	}
	return json.Marshal(v)
}

func (jsonFormat) SerializeString(v any) (string, error) {
	if v == nil {
		return "null", nil
	}
	if b, ok := v.([]byte); ok {
		return string(b), nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (jsonFormat) Deserialize(b []byte) (any, error) {
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return nil, fmt.Errorf("json wire format: %w", err)
	}
	return v, nil
}

func (jsonFormat) Name() string { return "json" }
