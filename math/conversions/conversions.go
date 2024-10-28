package conversions

import (
	"math"
)

const (
	ref1V   float64 = 1
	ref0775 float64 = 0.775
)

func Vpp2Vrms(vpp float64) float64 {
	return vpp / (2 * math.Sqrt(2))
}

func Vpp2dBV(vpp float64) float64 {
	return Vrms2dBV(Vpp2Vrms(vpp))
}

func Vpp2dBu(vpp float64) float64 {
	return Vrms2dBu(Vpp2Vrms(vpp))
}

func Vrms2dBV(vrms float64) float64 {
	return v2dB(vrms, ref1V)
}

func Vrms2dBu(vrms float64) float64 {
	return v2dB(vrms, ref0775)
}

func v2dB(v float64, ref float64) float64 {
	return 20 * math.Log10(v/ref)
}
