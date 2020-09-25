package compatnum

import (
	"bytes"
	"encoding/json"
	"strconv"
)

var (
	// MaxInt defines the maximum number that will be encoded as a JSON number rather than a string.
	// It is not save to change this value while a struct containing compatnum.Int or compatnum.Uint is being marshaled or unmarshaled.
	MaxInt = 0xFFFFFFFF
)

// Int represents a signed integer that can be expressed as both a number and a string in JSON.
type Int int64

// UnmarshalJSON implements the json.Unmarshaler interface and decodes both numbers and strings into int64s.
func (i *Int) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	// By convention, to approximate the behavior of Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
	// Source: https://golang.org/pkg/encoding/json/#Unmarshaler
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str := ""
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		// this is a string
		str = string(b[1 : len(b)-1])
	} else {
		str = string(b)
	}

	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}

	*i = Int(i64)
	return nil
}

// MarshalJSON implements the json.Marshaler interface and encodes int64s as numbers or strings depending on MaxInt.
func (i Int) MarshalJSON() ([]byte, error) {
	if i <= Int(MaxInt) {
		return json.Marshal(int64(i))
	}

	str := strconv.FormatInt(int64(i), 10)
	return json.Marshal(str)
}

// Uint represents an unsigned integer that can be expressed as both a number and a string in JSON.
type Uint uint64

// UnmarshalJSON implements the json.Unmarshaler interface and decodes both numbers and strings into uint64s.
func (u *Uint) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	// By convention, to approximate the behavior of Unmarshal itself, Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
	// Source: https://golang.org/pkg/encoding/json/#Unmarshaler
	if bytes.Equal(b, []byte("null")) {
		return nil
	}

	str := ""
	if len(b) > 2 && b[0] == '"' && b[len(b)-1] == '"' {
		// this is a string
		str = string(b[1 : len(b)-1])
	} else {
		str = string(b)
	}

	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}

	*u = Uint(u64)
	return nil
}

// MarshalJSON implements the json.Marshaler interface and encodes uint64s as numbers or strings depending on MaxInt.
func (u Uint) MarshalJSON() ([]byte, error) {
	if u <= Uint(MaxInt) {
		return json.Marshal(int64(u))
	}

	str := strconv.FormatUint(uint64(u), 10)
	return json.Marshal(str)
}
