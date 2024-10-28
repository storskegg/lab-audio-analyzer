// lpsd is a package for calculating the Power Spectral Density (PSD) of a given
// signal. It is a port of the matlab implementation found here:
// https://github.com/tobin/lpsd/blob/master/lpsd.m
// ...which itself is an implementation based on an academic paper by Michael
// Tröbs and Gerhard Heinzel, which you can find here:
// https://pure.mpg.de/pubman/faces/ViewItemOverviewPage.jsp?itemId=item_150688_1

package lpsd

// Lpsd calculates the Power Spectral Density (PSD) of a given signal utilizing
// a logarithmic frequency axis.
//func Lpsd(x, windowfcn, fmin, fmax, Jdes, Kdes, Kmin, fs, xi) (Pxx []float64, f []float64, C any, err error) {
//
//}
