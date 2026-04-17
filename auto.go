package wire

import "encoding/json"

type autoFormat struct{}

func (autoFormat) Serialize(v any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	switch val := v.(type) {
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	default:
		return json.Marshal(v)
	}
}

func (autoFormat) SerializeString(v any) (string, error) {
	if v == nil {
		return "", nil
	}
	switch val := v.(type) {
	case []byte:
		return string(val), nil
	case string:
		return val, nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}

func (autoFormat) Deserialize(b []byte) (any, error) {
	s := string(b)
	if !looksLikeJSON(s) {
		return s, nil
	}

	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		// JSON detection was a false positive (e.g. "2026-04-14"
		// starts with a digit but isn't valid JSON). Fall back to string.
		return s, nil
	}
	return v, nil
}

func (autoFormat) Name() string { return "auto" }
