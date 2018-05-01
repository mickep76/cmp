package cmp

import (
	"testing"
	"time"
)

type testPair struct {
	a      interface{}
	b      interface{}
	expEq  bool
	expNeq bool
	expLt  bool
	expGt  bool
	expLte bool
	expGte bool
}

type testInterface struct {
	Number int
}

func (t testInterface) Eq(b interface{}) (bool, error) {
	return t.Number == b.(testInterface).Number, nil
}

func (t testInterface) Lt(b interface{}) (bool, error) {
	return t.Number < b.(testInterface).Number, nil
}

var now = time.Now()
var testData = []testPair{
	testPair{a: true, b: true, expEq: true, expNeq: false},
	testPair{a: true, b: false, expEq: false, expNeq: true},

	testPair{a: int(1), b: int(1), expEq: true, expNeq: false},
	testPair{a: int(1), b: int(2), expEq: false, expNeq: true},

	testPair{a: int8(1), b: int8(1), expEq: true, expNeq: false},
	testPair{a: int8(1), b: int8(2), expEq: false, expNeq: true},

	testPair{a: int16(1), b: int16(1), expEq: true, expNeq: false},
	testPair{a: int16(1), b: int16(2), expEq: false, expNeq: true},

	testPair{a: int32(1), b: int32(1), expEq: true, expNeq: false},
	testPair{a: int32(1), b: int32(2), expEq: false, expNeq: true},

	testPair{a: int64(1), b: int64(1), expEq: true, expNeq: false},
	testPair{a: int64(1), b: int64(2), expEq: false, expNeq: true},

	testPair{a: uint(1), b: uint(1), expEq: true, expNeq: false},
	testPair{a: uint(1), b: uint(2), expEq: false, expNeq: true},

	testPair{a: uint8(1), b: uint8(1), expEq: true, expNeq: false},
	testPair{a: uint8(1), b: uint8(2), expEq: false, expNeq: true},

	testPair{a: uint16(1), b: uint16(1), expEq: true, expNeq: false},
	testPair{a: uint16(1), b: uint16(2), expEq: false, expNeq: true},

	testPair{a: uint32(1), b: uint32(1), expEq: true, expNeq: false},
	testPair{a: uint32(1), b: uint32(2), expEq: false, expNeq: true},

	testPair{a: uint64(1), b: uint64(1), expEq: true, expNeq: false},
	testPair{a: uint64(1), b: uint64(2), expEq: false, expNeq: true},

	testPair{a: float32(1), b: float32(1), expEq: true, expNeq: false},
	testPair{a: float32(1), b: float32(2), expEq: false, expNeq: true},

	testPair{a: float64(1), b: float64(1), expEq: true, expNeq: false},
	testPair{a: float64(1), b: float64(2), expEq: false, expNeq: true},

	testPair{a: now, b: now, expEq: true, expNeq: false},
	testPair{a: now, b: now.Add(time.Second), expEq: false, expNeq: true},

	testPair{a: "abc123", b: "abc123", expEq: true, expNeq: false},
	testPair{a: "abc123", b: "def456", expEq: false, expNeq: true},

	testPair{a: testInterface{1}, b: testInterface{1}, expEq: true, expNeq: false},
	testPair{a: testInterface{1}, b: testInterface{2}, expEq: false, expNeq: true},
}

func TestEq(t *testing.T) {
	_, err := Eq(1, "abc123")
	if err != ErrNotSameKind {
		t.Error(err)
	}

	_, err = Eq(testData, testData)
	if err != ErrKindNotSupported {
		t.Error(err)
	}

	for _, d := range testData {
		m, err := Eq(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expEq {
			t.Errorf("want:\n%v, got:\n%v", d.expEq, m)
		}

		m, err = Neq(d.a, d.b)
		if err != nil {
			t.Error(err)
		}
		if m != d.expNeq {
			t.Errorf("want:\n%v, got:\n%v", d.expNeq, m)
		}
	}
}