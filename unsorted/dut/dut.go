package dut

type Type int

const (
	TypeUnknown Type = iota
	TypeFilter
	TypeAmplifier
)

type DeviceUnderTest struct {
	Name string
	Type Type
}
