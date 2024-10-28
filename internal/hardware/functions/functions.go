package functions

type SupportsFunction int

const (
	SupportsFunctionNone SupportsFunction = iota << 1
	SupportsFunctionVrms
	SupportsFunctionIrms
	SupportsFunctionDDS
)
