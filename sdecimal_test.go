package sdecimal

import (
	"math/big"
	"testing"
)

func TestToRawBigInt(t *testing.T) {
	raw := big.NewInt(0).SetUint64(18123e15)
	sd := NewFromBigInt(raw, 18).ToRawBigInt()
	if sd.Cmp(raw) != 0 {
		t.Error("TestToRawBigInt", sd.String(), raw.String())
	}
}

func TestToRawBigFloat(t *testing.T) {
	raw := big.NewInt(0).SetUint64(18123e15)
	sd := NewFromBigInt(raw, 18).ToRawBigFloat()
	if sd.Cmp(big.NewFloat(0).SetInt(raw)) != 0 {
		t.Error("TestToRawBigFloat", sd.String(), raw.String())
	}
}

func TestToFloat(t *testing.T) {
	raw := 18.123
	sd := NewFromFloat(raw)
	if sd.MustToFloat() != raw {
		t.Error("TestToFloat", sd.String())
	}
}

func TestToGwei(t *testing.T) {
	raw := 18.123
	rawEther := raw / float64(1e9)
	bigRaw := big.NewInt(0).SetUint64(18.123 * 1e9)

	sd := NewFromEtherGwei(raw)
	if sd.MustToFloat() != rawEther {
		t.Error("TestToGwei-0", sd.String(), bigRaw.String())
	}

	if sd.ToGwei() != raw {
		t.Error("TestToGwei-1", sd.ToGwei(), bigRaw.String())
	}

	if sd.ToRawBigInt().Cmp(bigRaw) != 0 {
		t.Error("TestToGwei-2", sd.String(), bigRaw.String())
	}
}

func TestWithDecimals(t *testing.T) {
	raw := 18.123
	raw3Decimals := big.NewInt(0).SetUint64(18.123 * 1e3)
	raw10Decimals := big.NewInt(0).SetUint64(18.123 * 1e10)

	sd := NewFromFloat(raw)
	if sd.ToRawBigInt().Cmp(raw3Decimals) != 0 {
		t.Error("TestWithDecimals-0", sd.String(), raw3Decimals.String())
	}
	if sd.WithDecimals(10).ToRawBigInt().Cmp(raw10Decimals) != 0 {
		t.Error("TestWithDecimals-1", sd.String(), raw3Decimals.String())
	}

	if sd.WithDecimals(10).MustToFloat() != raw {
		t.Error("TestWithDecimals-2", sd.String())
	}
}

func TestNewFromFloatWithDecimals(t *testing.T) {
	raw := 18.123
	raw10Decimals := big.NewInt(0).SetUint64(18.123 * 1e10)
	sd := NewFromFloatWithDecimals(raw, 10)
	if sd.MustToFloat() != raw {
		t.Error("TestNewFromFloatWithDecimals-0", sd.String())
	}

	if sd.ToRawBigInt().Cmp(raw10Decimals) != 0 {
		t.Error("TestNewFromFloatWithDecimals-1", sd.String(), raw10Decimals.String())
	}
}

