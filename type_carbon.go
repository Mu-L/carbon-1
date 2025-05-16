package carbon

import (
	"bytes"
	"database/sql/driver"
)

// Scan implements driver.Scanner interface for Carbon struct.
func (c *Carbon) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		*c = *Parse(string(v), DefaultTimezone)
	case string:
		*c = *Parse(v, DefaultTimezone)
	case int64:
		*c = *CreateFromTimestamp(v, DefaultTimezone)
	case StdTime:
		*c = *CreateFromStdTime(v, DefaultTimezone)
	case *StdTime:
		*c = *CreateFromStdTime(*v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	return c.Error
}

// Value implements driver.Valuer interface for Carbon struct.
func (c Carbon) Value() (driver.Value, error) {
	if c.IsNil() || c.IsZero() || c.IsEmpty() {
		return nil, nil
	}
	if c.HasError() {
		return nil, c.Error
	}
	return c.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for Carbon struct.
func (c Carbon) MarshalJSON() ([]byte, error) {
	if c.IsNil() || c.IsZero() || c.IsEmpty() {
		return []byte(`null`), nil
	}
	if c.HasError() {
		return []byte(`null`), c.Error
	}
	v := c.Layout(DefaultLayout)
	b := make([]byte, 0, len(v)+2)
	b = append(b, '"')
	b = append(b, v...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements json.Unmarshal interface for Carbon struct.
func (c *Carbon) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" || v == "null" {
		return nil
	}
	*c = *ParseByLayout(v, DefaultLayout)
	return c.Error
}

// String implements the interface Stringer for Carbon struct.
func (c *Carbon) String() string {
	if c.IsInvalid() || c.IsZero() {
		return ""
	}
	return c.Layout(c.currentLayout)
}
