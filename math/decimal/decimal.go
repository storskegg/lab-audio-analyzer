// decimal is merely a wrapper around the govalues/decimal package, extending it
// with a Hash function that calculates a SHA-256_224 hash of the decimal's
// binary marshal representation.

package decimal

import (
	"crypto/sha256"
	"database/sql/driver"
	"fmt"

	"github.com/govalues/decimal"
)

type Hash [28]byte

type Decimal interface {
	Hash() (Hash, error)

	Zero() decimal.Decimal
	One() decimal.Decimal
	ULP() decimal.Decimal
	String() string
	Float64() (f float64, ok bool)
	Int64(scale int) (whole int64, frac int64, ok bool)
	MarshalText() ([]byte, error)
	MarshalBinary() ([]byte, error)
	Value() (driver.Value, error)
	Format(state fmt.State, verb rune)
	Prec() int
	Coef() uint64
	Scale() int
	MinScale() int
	IsInt() bool
	IsOne() bool
	WithinOne() bool
	Round(scale int) decimal.Decimal
	Pad(scale int) decimal.Decimal
	Rescale(scale int) decimal.Decimal
	Quantize(e decimal.Decimal) decimal.Decimal
	SameScale(e decimal.Decimal) bool
	Trunc(scale int) decimal.Decimal
	Trim(scale int) decimal.Decimal
	Ceil(scale int) decimal.Decimal
	Floor(scale int) decimal.Decimal
	Neg() decimal.Decimal
	Abs() decimal.Decimal
	CopySign(e decimal.Decimal) decimal.Decimal
	Sign() int
	IsPos() bool
	IsNeg() bool
	IsZero() bool
	Mul(e decimal.Decimal) (decimal.Decimal, error)
	MulExact(e decimal.Decimal, scale int) (decimal.Decimal, error)
	Pow(power int) (decimal.Decimal, error)
	PowInt(power int) (decimal.Decimal, error)
	Sqrt() (decimal.Decimal, error)
	Exp() (decimal.Decimal, error)
	Log() (decimal.Decimal, error)
	SubAbs(e decimal.Decimal) (decimal.Decimal, error)
	Sub(e decimal.Decimal) (decimal.Decimal, error)
	SubExact(e decimal.Decimal, scale int) (decimal.Decimal, error)
	Add(e decimal.Decimal) (decimal.Decimal, error)
	AddExact(e decimal.Decimal, scale int) (decimal.Decimal, error)
	FMA(e decimal.Decimal, f decimal.Decimal) (decimal.Decimal, error)
	FMAExact(e decimal.Decimal, f decimal.Decimal, scale int) (decimal.Decimal, error)
	SubMul(e decimal.Decimal, f decimal.Decimal) (decimal.Decimal, error)
	SubMulExact(e decimal.Decimal, f decimal.Decimal, scale int) (decimal.Decimal, error)
	AddMul(e decimal.Decimal, f decimal.Decimal) (decimal.Decimal, error)
	AddMulExact(e decimal.Decimal, f decimal.Decimal, scale int) (decimal.Decimal, error)
	SubQuo(e decimal.Decimal, f decimal.Decimal) (decimal.Decimal, error)
	SubQuoExact(e decimal.Decimal, f decimal.Decimal, scale int) (decimal.Decimal, error)
	AddQuo(e decimal.Decimal, f decimal.Decimal) (decimal.Decimal, error)
	AddQuoExact(e decimal.Decimal, f decimal.Decimal, scale int) (decimal.Decimal, error)
	Inv() (decimal.Decimal, error)
	Quo(e decimal.Decimal) (decimal.Decimal, error)
	QuoExact(e decimal.Decimal, scale int) (decimal.Decimal, error)
	QuoRem(e decimal.Decimal) (q decimal.Decimal, r decimal.Decimal, err error)
	Max(e decimal.Decimal) decimal.Decimal
	Min(e decimal.Decimal) decimal.Decimal
	Clamp(min decimal.Decimal, max decimal.Decimal) (decimal.Decimal, error)
	CmpTotal(e decimal.Decimal) int
	CmpAbs(e decimal.Decimal) int
	Equal(e decimal.Decimal) bool
	Less(e decimal.Decimal) bool
	Cmp(e decimal.Decimal) int
}

type dec struct {
	decimal.Decimal
}

func New(v float64) (Decimal, error) {
	d, err := decimal.NewFromFloat64(v)
	if err != nil {
		return nil, err
	}
	dd := &dec{d}
	return dd, nil
}

func (d *dec) Hash() (Hash, error) {
	v, err := d.MarshalBinary()
	if err != nil {
		return zero(), err
	}

	return sha256.Sum224(v), nil
}

func zero() Hash {
	h := Hash{}

	for i, _ := range h {
		h[i] = 0
	}

	return h
}
