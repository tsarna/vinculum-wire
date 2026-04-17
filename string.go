package wire

import "fmt"

type stringFormat struct{}

func (stringFormat) Serialize(v any) ([]byte, error) {
	s, err := scalarToString(v)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (stringFormat) SerializeString(v any) (string, error) {
	return scalarToString(v)
}

func (stringFormat) Deserialize(b []byte) (any, error) {
	return string(b), nil
}

func (stringFormat) Name() string { return "string" }

// scalarToString converts scalar Go values to their string representation.
// []byte passes through (as a string), strings pass through, numbers and
// bools use their canonical form. Objects, lists, and nil are errors.
func scalarToString(v any) (string, error) {
	switch val := v.(type) {
	case []byte:
		return string(val), nil
	case string:
		return val, nil
	case bool:
		if val {
			return "true", nil
		}
		return "false", nil
	case int:
		return fmt.Sprintf("%d", val), nil
	case int8:
		return fmt.Sprintf("%d", val), nil
	case int16:
		return fmt.Sprintf("%d", val), nil
	case int32:
		return fmt.Sprintf("%d", val), nil
	case int64:
		return fmt.Sprintf("%d", val), nil
	case uint:
		return fmt.Sprintf("%d", val), nil
	case uint8:
		return fmt.Sprintf("%d", val), nil
	case uint16:
		return fmt.Sprintf("%d", val), nil
	case uint32:
		return fmt.Sprintf("%d", val), nil
	case uint64:
		return fmt.Sprintf("%d", val), nil
	case float32:
		return fmt.Sprintf("%g", val), nil
	case float64:
		return fmt.Sprintf("%g", val), nil
	case nil:
		return "", fmt.Errorf("string wire format: cannot serialize nil")
	default:
		return "", fmt.Errorf("string wire format: cannot serialize %T", v)
	}
}
