package util

import "testing"

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d, wanted 1", got)
	}
	got = Abs(0)
	if got != 0 {
		t.Errorf("Abs(0) = %d, wanted 0", got)
	}

	got = Abs(1)
	if got != 1 {
		t.Errorf("Abs(1) = %d, wanted 1", got)
	}

	got = Abs(-9223372036854775808)
	if got != -9223372036854775808 {
		t.Errorf("Abs(-9223372036854775808) = %d, wanted -9223372036854775808", got)
	}

}

func TestMax(t *testing.T) {
	got := Max(3, 5)
	if got != 5 {
		t.Errorf("Max(3,5) = %d, wanted 5", got)
	}
	got = Max(-3, -5)
	if got != -3 {
		t.Errorf("Max(-3,-5) = %d, wanted -3", got)
	}
	got = Max(0, -5)
	if got != -0 {
		t.Errorf("Max(0,-5) = %d, wanted 0", got)
	}
	got = Max(-9223372036854775808, 9223372036854775807)
	if got != 9223372036854775807 {
		t.Errorf("Max(-9223372036854775808,9223372036854775807) = %d, wanted 9223372036854775807", got)
	}

}

func TestMin(t *testing.T) {
	got := Min(3, 5)
	if got != 3 {
		t.Errorf("Min(3,5) = %d, wanted 3", got)
	}
	got = Min(-3, -5)
	if got != -5 {
		t.Errorf("Max(-3,-5) = %d, wanted -5", got)
	}
	got = Min(0, -5)
	if got != -5 {
		t.Errorf("Max(0,-5) = %d, wanted -5", got)
	}
	got = Min(-9223372036854775808, 9223372036854775807)
	if got != -9223372036854775808 {
		t.Errorf("Max(-9223372036854775808,9223372036854775807) = %d, wanted -9223372036854775808", got)
	}
	got = Min(5, -5)
	if got != -5 {
		t.Errorf("Max(5,-5) = %d, wanted -5", got)
	}

}
