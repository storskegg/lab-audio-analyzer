package binning

import (
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
	Xi   float64 = 0.3 // Fractional segment overlap (ξ); 0 <= Xi <= 1
	Kdes         = 100 // Number of desired averages
	Kmin         = 3   // Minimimum number of averages
	Fs           = 1   // Sampling frequency of a time series
	N            = 100 // Number of data; TODO
	// ravg is the smallest frequency resolution with Kdes averages
	ravg = (Fs / float64(N)) * (1 + (1-Xi)*(Kdes-1))
	// rmin is the frequency resolution corresponding to Kmin averages
	rmin = (Fs / float64(N)) * (1 + (1-Xi)*(Kmin-1))
)

type Bins interface {
	FreqMin() float64
	FreqEnd() float64

	G() float64 // Table 1 (g)
}

type bins struct {
	mu sync.RWMutex

	Fmin float64
	Fmax float64

	J    int // Fourier frequencies
	Jdes int // Desired number of Fourier frequencies

	bins []Bin
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

// ravg returns the smallest frequency resolution with Kdes averages.
//func ravg(N int) float64 {
//	return (Fs / float64(N)) * (1 + (1-Xi)*(Kdes-1))
//}

// rpp Fn(18) of the pdf
func rpp(idx int, fmin float64, fmax float64, lenBins int) float64 {
	rpi := rp(idx, fmin, fmax, lenBins)

	if rpi >= ravg {
		return rpi
	} else if xxx := math.Sqrt(ravg * rpi); rpi < ravg && xxx > rmin {
		return xxx
	}

	return rmin
}
