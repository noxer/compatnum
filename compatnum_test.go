package compatnum_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/noxer/compatnum"
)

type testStructInt struct {
	MyInt compatnum.Int
}

func TestIntUnmarshalInt(t *testing.T) {
	data := []byte(`{"myint":123}`)
	ts := testStructInt{}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != 123 {
		t.Errorf("MyInt was %d, expected %d", ts.MyInt, 123)
	}
}

func TestIntUnmarshalNegInt(t *testing.T) {
	data := []byte(`{"myint":-123}`)
	ts := testStructInt{}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != -123 {
		t.Errorf("MyInt was %d, expected %d", ts.MyInt, -123)
	}
}

func TestIntUnmarshalStr(t *testing.T) {
	data := []byte(`{"myint":"123"}`)
	ts := testStructInt{}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != 123 {
		t.Errorf("MyInt was %d, expected %d", ts.MyInt, 123)
	}
}

func TestIntUnmarshalNull(t *testing.T) {
	data := []byte(`{"myint":null}`)
	ts := testStructInt{MyInt: 42}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != 42 {
		t.Errorf("MyInt was %d, expected %d (unchanged)", ts.MyInt, 42)
	}
}

func TestIntUnmarshalErr(t *testing.T) {
	data := []byte(`{"myint":"hello"}`)
	ts := testStructInt{}
	err := json.Unmarshal(data, &ts)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestIntUnmarshalEmpty(t *testing.T) {
	ts := testStructInt{MyInt: 42}
	err := ts.MyInt.UnmarshalJSON([]byte{})
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != 42 {
		t.Errorf("MyInt was %d, expected %d (unchanged)", ts.MyInt, 42)
	}
}

func TestIntUnmarshalNil(t *testing.T) {
	ts := testStructInt{MyInt: 42}
	err := ts.MyInt.UnmarshalJSON(nil)
	if err != nil {
		t.Error(err)
	}
	if ts.MyInt != 42 {
		t.Errorf("MyInt was %d, expected %d (unchanged)", ts.MyInt, 42)
	}
}

func TestIntMarshalInt(t *testing.T) {
	i := compatnum.Int(42)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("42")) {
		t.Errorf("Output was %s, expected %s", p, "42")
	}
}

func TestIntMarshalBigInt(t *testing.T) {
	i := compatnum.Int(8589934591)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("\"8589934591\"")) {
		t.Errorf("Output was %s, expected %s", p, "\"8589934591\"")
	}
}

func TestIntMarshalSetMax(t *testing.T) {
	compatnum.MaxInt = 10
	t.Cleanup(func() {
		compatnum.MaxInt = 0xFFFFFFFF
	})

	i := compatnum.Int(10)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("10")) {
		t.Errorf("Output was %s, expected %s", p, "10")
	}

	i = compatnum.Int(11)
	p, err = json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("\"11\"")) {
		t.Errorf("Output was %s, expected %s", p, "\"11\"")
	}
}

type testStructUint struct {
	MyUint compatnum.Uint
}

func TestUintUnmarshalUint(t *testing.T) {
	data := []byte(`{"myuint":123}`)
	ts := testStructUint{}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyUint != 123 {
		t.Errorf("MyUint was %d, expected %d", ts.MyUint, 123)
	}
}

func TestUintUnmarshalNegUint(t *testing.T) {
	data := []byte(`{"myuint":-123}`)
	ts := testStructUint{}
	err := json.Unmarshal(data, &ts)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestUintUnmarshalStr(t *testing.T) {
	data := []byte(`{"myuint":"123"}`)
	ts := testStructUint{}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyUint != 123 {
		t.Errorf("MyUint was %d, expected %d", ts.MyUint, 123)
	}
}

func TestUintUnmarshalNull(t *testing.T) {
	data := []byte(`{"myuint":null}`)
	ts := testStructUint{MyUint: 42}
	err := json.Unmarshal(data, &ts)
	if err != nil {
		t.Error(err)
	}
	if ts.MyUint != 42 {
		t.Errorf("MyUint was %d, expected %d (unchanged)", ts.MyUint, 42)
	}
}

func TestUintUnmarshalErr(t *testing.T) {
	data := []byte(`{"myuint":"hello"}`)
	ts := testStructUint{}
	err := json.Unmarshal(data, &ts)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestUintUnmarshalEmpty(t *testing.T) {
	ts := testStructUint{MyUint: 42}
	err := ts.MyUint.UnmarshalJSON([]byte{})
	if err != nil {
		t.Error(err)
	}
	if ts.MyUint != 42 {
		t.Errorf("MyUint was %d, expected %d (unchanged)", ts.MyUint, 42)
	}
}

func TestUintUnmarshalNil(t *testing.T) {
	ts := testStructUint{MyUint: 42}
	err := ts.MyUint.UnmarshalJSON(nil)
	if err != nil {
		t.Error(err)
	}
	if ts.MyUint != 42 {
		t.Errorf("MyUint was %d, expected %d (unchanged)", ts.MyUint, 42)
	}
}

func TestUintMarshalUint(t *testing.T) {
	i := compatnum.Uint(42)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("42")) {
		t.Errorf("Output was %s, expected %s", p, "42")
	}
}

func TestUintMarshalBigUint(t *testing.T) {
	i := compatnum.Uint(8589934591)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("\"8589934591\"")) {
		t.Errorf("Output was %s, expected %s", p, "\"8589934591\"")
	}
}

func TestUintMarshalSetMax(t *testing.T) {
	compatnum.MaxInt = 10
	t.Cleanup(func() {
		compatnum.MaxInt = 0xFFFFFFFF
	})

	i := compatnum.Uint(10)
	p, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("10")) {
		t.Errorf("Output was %s, expected %s", p, "10")
	}

	i = compatnum.Uint(11)
	p, err = json.Marshal(i)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(p, []byte("\"11\"")) {
		t.Errorf("Output was %s, expected %s", p, "\"11\"")
	}
}
