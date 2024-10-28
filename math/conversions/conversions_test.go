package conversions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVpp2Vrms(t *testing.T) {
	type args struct {
		vpp float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Vpp = 0.25",
			args: args{
				vpp: 0.25,
			},
			want: "0.09",
		},
		{
			name: "Vpp = 0.5",
			args: args{
				vpp: 0.5,
			},
			want: "0.18",
		},
		{
			name: "Vpp = 0.75",
			args: args{
				vpp: 0.75,
			},
			want: "0.27",
		},
		{
			name: "Vpp = 1.0",
			args: args{
				vpp: 1.0,
			},
			want: "0.35",
		},
		{
			name: "Vpp = 1.25",
			args: args{
				vpp: 1.25,
			},
			want: "0.44",
		},
		{
			name: "Vpp = 1.5",
			args: args{
				vpp: 1.5,
			},
			want: "0.53",
		},
		{
			name: "Vpp = 1.75",
			args: args{
				vpp: 1.75,
			},
			want: "0.62",
		},
		{
			name: "Vpp = 2.0",
			args: args{
				vpp: 2.0,
			},
			want: "0.71",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vrms := Vpp2Vrms(tt.args.vpp)
			// To avoid floating point precision issues, we round the result to
			// 2 decimal places, and string compare.
			got := fmt.Sprintf("%.2f", vrms)

			assert.Equal(t, got, tt.want, "Vpp2Vrms() = %v, want %v", got, tt.want)
		})
	}
}

func TestVpp2dBV(t *testing.T) {
	type args struct {
		vpp float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Vpp = 0.25",
			args: args{
				vpp: 0.25,
			},
			want: "-21.07",
		},
		{
			name: "Vpp = 0.5",
			args: args{
				vpp: 0.5,
			},
			want: "-15.05",
		},
		{
			name: "Vpp = 0.75",
			args: args{
				vpp: 0.75,
			},
			want: "-11.53",
		},
		{
			name: "Vpp = 1.0",
			args: args{
				vpp: 1.0,
			},
			want: "-9.03",
		},
		{
			name: "Vpp = 2",
			args: args{
				vpp: 2.0,
			},
			want: "-3.01",
		},
		{
			name: "Vpp = 3",
			args: args{
				vpp: 3.0,
			},
			want: "0.51",
		},
		{
			name: "Vpp = 4",
			args: args{
				vpp: 4.0,
			},
			want: "3.01",
		},
		{
			name: "Vpp = 5",
			args: args{
				vpp: 5.0,
			},
			want: "4.95",
		},
		{
			name: "Vpp = 10",
			args: args{
				vpp: 10.0,
			},
			want: "10.97",
		},
		{
			name: "Vpp = 15",
			args: args{
				vpp: 15.0,
			},
			want: "14.49",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbv := Vpp2dBV(tt.args.vpp)
			got := fmt.Sprintf("%.2f", dbv)

			assert.Equal(t, tt.want, got, "Vpp2dBV() = %v, want %v", got, tt.want)
		})
	}
}

func TestVpp2dBu(t *testing.T) {
	type args struct {
		vpp float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Vpp = 0.25",
			args: args{
				vpp: 0.25,
			},
			want: "-18.86",
		},
		{
			name: "Vpp = 0.5",
			args: args{
				vpp: 0.5,
			},
			want: "-12.84",
		},
		{
			name: "Vpp = 0.75",
			args: args{
				vpp: 0.75,
			},
			want: "-9.32",
		},
		{
			name: "Vpp = 1.0",
			args: args{
				vpp: 1.0,
			},
			want: "-6.82",
		},
		{
			name: "Vpp = 2",
			args: args{
				vpp: 2.0,
			},
			want: "-0.80",
		},
		{
			name: "Vpp = 3",
			args: args{
				vpp: 3.0,
			},
			want: "2.73",
		},
		{
			name: "Vpp = 4",
			args: args{
				vpp: 4.0,
			},
			want: "5.22",
		},
		{
			name: "Vpp = 5",
			args: args{
				vpp: 5.0,
			},
			want: "7.16",
		},
		{
			name: "Vpp = 10",
			args: args{
				vpp: 10.0,
			},
			want: "13.18",
		},
		{
			name: "Vpp = 15",
			args: args{
				vpp: 15.0,
			},
			want: "16.70",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbu := Vpp2dBu(tt.args.vpp)
			got := fmt.Sprintf("%.2f", dbu)

			assert.Equal(t, tt.want, got, "Vpp2dBu() = %v, want %v", got, tt.want)
		})
	}
}

func TestVrms2dBV(t *testing.T) {
	type args struct {
		vrms float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Vrms = 0.25",
			args: args{
				vrms: 0.25,
			},
			want: "-12.04",
		},
		{
			name: "Vrms = 0.5",
			args: args{
				vrms: 0.5,
			},
			want: "-6.02",
		},
		{
			name: "Vrms = 0.75",
			args: args{
				vrms: 0.75,
			},
			want: "-2.50",
		},
		{
			name: "Vrms = 1.0",
			args: args{
				vrms: 1.0,
			},
			want: "0.00",
		},

		{
			name: "Vrms = 1.5",
			args: args{
				vrms: 1.5,
			},
			want: "3.52",
		},
		{
			name: "Vrms = 2.0",
			args: args{
				vrms: 2.0,
			},
			want: "6.02",
		},
		{
			name: "Vrms = 3.0",
			args: args{
				vrms: 3.0,
			},
			want: "9.54",
		},
		{
			name: "Vrms = 4.0",
			args: args{
				vrms: 4.0,
			},
			want: "12.04",
		},
		{
			name: "Vrms = 5.0",
			args: args{
				vrms: 5.0,
			},
			want: "13.98",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbv := Vrms2dBV(tt.args.vrms)
			got := fmt.Sprintf("%.2f", dbv)

			assert.Equal(t, tt.want, got, "Vrms2dBV() = %v, want %v", got, tt.want)
		})
	}
}

func TestVrms2dBu(t *testing.T) {
	type args struct {
		vrms float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Vrms = 0.25",
			args: args{
				vrms: 0.25,
			},
			want: "-9.83",
		},
		{
			name: "Vrms = 0.5",
			args: args{
				vrms: 0.5,
			},
			want: "-3.81",
		},
		{
			name: "Vrms = 0.75",
			args: args{
				vrms: 0.75,
			},
			want: "-0.28",
		},
		{
			name: "Vrms = 1.0",
			args: args{
				vrms: 1.0,
			},
			want: "2.21",
		},

		{
			name: "Vrms = 1.5",
			args: args{
				vrms: 1.5,
			},
			want: "5.74",
		},
		{
			name: "Vrms = 2.0",
			args: args{
				vrms: 2.0,
			},
			want: "8.23",
		},
		{
			name: "Vrms = 3.0",
			args: args{
				vrms: 3.0,
			},
			want: "11.76",
		},
		{
			name: "Vrms = 4.0",
			args: args{
				vrms: 4.0,
			},
			want: "14.26",
		},
		{
			name: "Vrms = 5.0",
			args: args{
				vrms: 5.0,
			},
			want: "16.19",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbu := Vrms2dBu(tt.args.vrms)
			got := fmt.Sprintf("%.2f", dbu)

			assert.Equal(t, tt.want, got, "Vrms2dBu() = %v, want %v", got, tt.want)
		})
	}
}

func Test_v2dB(t *testing.T) {
	t.Skip("testing of this method is thoroughly covered by other tests")
}
