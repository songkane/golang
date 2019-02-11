// Package cmath unit test
// Created by chenguolin 2019-02-11
package math

import "testing"

func TestMaxInt(t *testing.T) {
	x := int(1)
	y := int(2)

	if MaxInt(x, y) != y {
		t.Fatal("TestMaxInt failed ~")
	}

	if MaxInt(y, x) != y {
		t.Fatal("TestMaxInt failed ~")
	}
}

func TestMinInt(t *testing.T) {
	x := int(1)
	y := int(2)

	if MinInt(x, y) != x {
		t.Fatal("TestMinInt failed ~")
	}

	if MinInt(y, x) != x {
		t.Fatal("TestMinInt failed ~")
	}
}

func TestMaxInt8(t *testing.T) {
	x := int8(1)
	y := int8(2)

	if MaxInt8(x, y) != y {
		t.Fatal("TestMaxInt8 failed ~")
	}

	if MaxInt8(y, x) != y {
		t.Fatal("TestMaxInt8 failed ~")
	}
}

func TestMinInt8(t *testing.T) {
	x := int8(1)
	y := int8(2)

	if MinInt8(x, y) != x {
		t.Fatal("TestMinInt8 failed ~")
	}

	if MinInt8(y, x) != x {
		t.Fatal("TestMinInt8 failed ~")
	}
}

func TestMaxInt16(t *testing.T) {
	x := int16(1)
	y := int16(2)

	if MaxInt16(x, y) != y {
		t.Fatal("TestMaxInt16 failed ~")
	}

	if MaxInt16(y, x) != y {
		t.Fatal("TestMaxInt16 failed ~")
	}
}

func TestMinInt16(t *testing.T) {
	x := int16(1)
	y := int16(2)

	if MinInt16(x, y) != x {
		t.Fatal("TestMinInt16 failed ~")
	}

	if MinInt16(y, x) != x {
		t.Fatal("TestMinInt16 failed ~")
	}
}

func TestMaxInt32(t *testing.T) {
	x := int32(1)
	y := int32(2)

	if MaxInt32(x, y) != y {
		t.Fatal("TestMaxInt32 failed ~")
	}

	if MaxInt32(y, x) != y {
		t.Fatal("TestMaxInt32 failed ~")
	}
}

func TestMinInt32(t *testing.T) {
	x := int32(1)
	y := int32(2)

	if MinInt32(x, y) != x {
		t.Fatal("TestMinInt32 failed ~")
	}

	if MinInt32(y, x) != x {
		t.Fatal("TestMinInt32 failed ~")
	}
}

func TestMaxInt64(t *testing.T) {
	x := int64(1)
	y := int64(2)

	if MaxInt64(x, y) != y {
		t.Fatal("TestMaxInt64 failed ~")
	}

	if MaxInt64(y, x) != y {
		t.Fatal("TestMaxInt64 failed ~")
	}
}

func TestMinInt64(t *testing.T) {
	x := int64(1)
	y := int64(2)

	if MinInt64(x, y) != x {
		t.Fatal("TestMinInt64 failed ~")
	}

	if MinInt64(y, x) != x {
		t.Fatal("TestMinInt64 failed ~")
	}
}

func TestMaxFloat32(t *testing.T) {
	x := float32(1.0)
	y := float32(2.0)

	if MaxFloat32(x, y) != y {
		t.Fatal("TestMaxFloat32 failed ~")
	}

	if MaxFloat32(y, x) != y {
		t.Fatal("TestMaxFloat32 failed ~")
	}
}

func TestMinFloat32(t *testing.T) {
	x := float32(1)
	y := float32(2)

	if MinFloat32(x, y) != x {
		t.Fatal("TestMinFloat32 failed ~")
	}

	if MinFloat32(y, x) != x {
		t.Fatal("TestMinFloat32 failed ~")
	}
}

func TestMaxFloat64(t *testing.T) {
	x := float64(1.0)
	y := float64(2.0)

	if MaxFloat64(x, y) != y {
		t.Fatal("TestMaxFloat64 failed ~")
	}

	if MaxFloat64(y, x) != y {
		t.Fatal("TestMaxFloat64 failed ~")
	}
}

func TestMinFloat64(t *testing.T) {
	x := float64(1)
	y := float64(2)

	if MinFloat64(x, y) != x {
		t.Fatal("TestMinFloat64 failed ~")
	}

	if MinFloat64(y, x) != x {
		t.Fatal("TestMinFloat64 failed ~")
	}
}
