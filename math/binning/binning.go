package binning

import (
	"errors"
	"math"
	"sync"
)

// Note that the naming in this package differs from idiomatic conventions, and
// instead attempts to remain close to the original paper's defined terms in
// Table 1 of https://hdl.handle.net/11858/00-001M-0000-0013-4BFD-C
//
// Where a variable or function is defined, and is not very semantic, I've
// noted the reference to Table 1 in the comments.

const (
	Fs = 1 // Sampling frequency of a time series
)

type Bins interface {
	FreqMin() float64
	FreqMax() float64
}

func NewBins(fmin float64, fmax float64, qtyBins int) (*bins, error) {
	if qtyBins%2 == 0 {
		return nil, errors.New("number of bins must be odd")
	}

	b := &bins{
		Fmin:    fmin,
		Fmax:    fmax,
		g:       g(fmin, fmax),
		NumBins: qtyBins,
		bins:    make([]Bin, qtyBins),
	}

	fNext := fmin
	var fp float64

	for i := 0; i < qtyBins; i++ {
		fp = rp(i, fmin, fmax, qtyBins)
		b.bins[i] = NewBin(fNext)
		fNext += fp
	}

	return b, nil
}

type bins struct {
	mu sync.RWMutex

	Fmin float64
	Fmax float64
	g    float64

	NumBins int

	bins []Bin
}

func (b *bins) FreqMin() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.Fmin
}

func (b *bins) FreqMax() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.Fmax
}

// g is an abbreviation
func g(min, max float64) float64 {
	return math.Log(max) - math.Log(min)
}

// f returns the Fourier frequency at index idx given the min and max
// frequencies of the set range, based on the WOSA method.
func f(idx int, fmin float64, fmax float64, lenBins int) float64 {
	return fmin * math.Exp((float64(idx)*g(fmin, fmax))/float64(lenBins-1))
}

// rp returns the difference between the given Fourier frequency and the next one.
func rp(idx int, fmin float64, fmax float64, lenBins int) float64 {
	return f(idx, fmin, fmax, lenBins) * (math.Exp(g(fmin, fmax)/float64(lenBins-1)) - 1)
}
