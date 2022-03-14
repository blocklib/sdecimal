package sdecimal

import (
	"math"
	"math/big"

	"github.com/shopspring/decimal"
)

type SDecimal struct {
	value decimal.Decimal
}

func New(v decimal.Decimal) SDecimal {
	ret := SDecimal{value: v}
	return ret
}

func NewFromBigInt(v *big.Int, decimals uint8) SDecimal {
	dv := decimal.NewFromBigInt(v, -int32(decimals))
	return New(dv)
}

func NewFromString(v string) (SDecimal, error) {
	dv, err := decimal.NewFromString(v)
	if err != nil {
		return SDecimal{}, err
	}
	return New(dv), nil
}

func ReqireFromString(v string) SDecimal {
	dv, _ := decimal.NewFromString(v)
	dv.Float64()
	return New(dv)
}

func NewFromInt(v int64) SDecimal {
	return New(decimal.NewFromInt(v))
}

func NewFromUInt(v uint64) SDecimal {
	return NewFromBigInt(big.NewInt(0).SetUint64(v), 0)
}

func NewFromFloat(v float64) SDecimal {
	return New(decimal.NewFromFloat(v))
}

func NewFromFloatWithDecimals(v float64, decimals uint8) SDecimal {
	var i, ex = big.NewInt(10), big.NewInt(int64(decimals))
	i.Exp(i, ex, nil)

	f := big.NewFloat(v)
	f.Mul(f, big.NewFloat(0).SetInt(i))

	ret, _ := f.Int(nil)
	return NewFromBigInt(ret, decimals)
}

func NewFromEtherKwei(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e3))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func NewFromEtherMwei(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e6))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func NewFromEtherGwei(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e9))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func NewFromMicroether(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e12))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func NewFromMilliether(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e15))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func NewFromEtherWei(v float64) SDecimal {
	fv := big.NewFloat(0).Mul(big.NewFloat(v), big.NewFloat(1e18))
	r, _ := fv.Int(nil)
	return NewFromBigInt(r, 18)
}

func (d SDecimal) WithDecimals(decimals uint8) SDecimal {
	var i, ex = big.NewInt(10), big.NewInt(int64(decimals))
	i.Exp(i, ex, nil)

	f := d.ToBigFloat()
	f.Mul(f, big.NewFloat(0).SetInt(i))

	ret, _ := f.Int(nil)
	return NewFromBigInt(ret, decimals)
}

func (d SDecimal) Value() decimal.Decimal {
	return d.value
}

//返回使用bigint表达的形势，即 v * 10^ex
func (d SDecimal) ToRawBigInt() *big.Int {
	f := d.value.BigFloat()
	exponent := int64(d.value.Exponent())
	//Decimal里v*10^ex ex是负数表示
	if d.value.Exponent() < 0 {
		exponent = int64(math.Abs(float64(exponent)))
	}
	var i, ex = big.NewInt(10), big.NewInt(exponent)
	i.Exp(i, ex, nil)
	if d.value.Exponent() < 0 {
		f.Mul(f, big.NewFloat(0).SetInt(i))
	} else {
		f.Quo(f, big.NewFloat(0).SetInt(i))
	}

	ret, _ := f.Int(nil)
	return ret
}

//只返回整数部分
func (d SDecimal) ToBigInt() *big.Int {
	return d.value.BigInt()
}

func (d SDecimal) ToBigFloat() *big.Float {
	return d.value.BigFloat()
}

//返回使用bigfloat表达的形势，即 v * 10^ex
func (d SDecimal) ToRawBigFloat() *big.Float {
	f := d.value.BigFloat()
	exponent := int64(d.value.Exponent())
	//Decimal里v*10^ex ex是负数表示
	if d.value.Exponent() < 0 {
		exponent = int64(math.Abs(float64(exponent)))
	}
	var i, ex = big.NewInt(10), big.NewInt(exponent)
	i.Exp(i, ex, nil)
	if d.value.Exponent() < 0 {
		f.Mul(f, big.NewFloat(0).SetInt(i))
	} else {
		f.Quo(f, big.NewFloat(0).SetInt(i))
	}

	return f
}

//返回最接近的float64值，exact表示是否精确表示
func (d SDecimal) ToFloat() (f float64, exact bool) {
	return d.value.Float64()
}

func (d SDecimal) MustToFloat() float64 {
	f, _ := d.value.Float64()
	return f
}

func (d SDecimal) ToString() string {
	return d.value.String()
}

func (d SDecimal) String() string {
	return d.value.String()
}

func (d SDecimal) ToGwei() float64 {
	f := d.ToRawBigFloat()
	var i, ex = big.NewInt(10), big.NewInt(int64(9))
	i.Exp(i, ex, nil)
	f.Quo(f, big.NewFloat(0).SetInt(i))

	v, _ := f.Float64()
	return v
}

func (d SDecimal) Add(other SDecimal) SDecimal {
	return New(d.value.Add(other.value))
}

func (d SDecimal) Sub(other SDecimal) SDecimal {
	return New(d.value.Sub(other.value))
}

func (d SDecimal) Mul(other SDecimal) SDecimal {
	return New(d.value.Mul(other.value))
}

func (d SDecimal) Div(other SDecimal) SDecimal {
	return New(d.value.Div(other.value))
}

func (d SDecimal) DivRound(other SDecimal, precision int32) SDecimal {
	return New(d.value.DivRound(other.value, precision))
}

func (d SDecimal) Mod(other SDecimal) SDecimal {
	return New(d.value.Mod(other.value))
}

func (d SDecimal) AddInt(d2 int64) SDecimal {
	return New(d.value.Add(decimal.NewFromInt(d2)))
}

func (d SDecimal) SubInt(d2 int64) SDecimal {
	return New(d.value.Sub(decimal.NewFromInt(d2)))
}

func (d SDecimal) MulInt(d2 int64) SDecimal {
	return New(d.value.Mul(decimal.NewFromInt(d2)))
}

func (d SDecimal) DivInt(d2 int64) SDecimal {
	return New(d.value.Div(decimal.NewFromInt(d2)))
}

func (d SDecimal) DivRoundInt(d2 int64, precision int32) SDecimal {
	return New(d.value.DivRound(decimal.NewFromInt(d2), precision))
}

func (d SDecimal) ModInt(d2 int64) SDecimal {
	return New(d.value.Mod(decimal.NewFromInt(d2)))
}

func (d SDecimal) AddFloat(d2 float64) SDecimal {
	return New(d.value.Add(decimal.NewFromFloat(d2)))
}

func (d SDecimal) SubFloat(d2 float64) SDecimal {
	return New(d.value.Sub(decimal.NewFromFloat(d2)))
}

func (d SDecimal) MulFloat(d2 float64) SDecimal {
	return New(d.value.Mul(decimal.NewFromFloat(d2)))
}

func (d SDecimal) DivFloat(d2 float64) SDecimal {
	return New(d.value.Div(decimal.NewFromFloat(d2)))
}

func (d SDecimal) DivRoundFloat(d2 float64, precision int32) SDecimal {
	return New(d.value.DivRound(decimal.NewFromFloat(d2), precision))
}

func (d SDecimal) ModFloat(d2 float64) SDecimal {
	return New(d.value.Mod(decimal.NewFromFloat(d2)))
}

func (d SDecimal) Com(other SDecimal) int {
	return d.value.Cmp(other.value)
}

// 保留指定小数位数
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
// 	   NewFromFloat(5.45).Round(1).String() // output: "5.5"
// 	   NewFromFloat(545).Round(-1).String() // output: "550"
func (d SDecimal) Round(precision int32) SDecimal {
	return New(d.value.Round(precision))
}

// 四舍五入到指定位数
// Examples:
//
// 	   NewFromFloat(5.45).Round(1).String() // output: "5.4"
// 	   NewFromFloat(545).Round(-1).String() // output: "540"
// 	   NewFromFloat(5.46).Round(1).String() // output: "5.5"
// 	   NewFromFloat(546).Round(-1).String() // output: "550"
// 	   NewFromFloat(5.55).Round(1).String() // output: "5.6"
// 	   NewFromFloat(555).Round(-1).String() // output: "560"
//
func (d SDecimal) RoundBank(precision int32) SDecimal {
	return New(d.value.Round(precision))
}

