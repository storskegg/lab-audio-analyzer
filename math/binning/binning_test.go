package binning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBins(t *testing.T) {
	type args struct {
		fmin    float64
		fmax    float64
		qtyBins int
	}
	tests := []struct {
		name    string
		args    args
		want    *bins
		wantErr bool
	}{
		{
			name: "Happy Path: small range, small bins",
			args: args{
				fmin:    10,
				fmax:    100,
				qtyBins: 11,
			},
			want: &bins{
				Fmin:    10,
				Fmax:    100,
				g:       2.302585092994046,
				NumBins: 11,
				binsByIndex: map[int]string{}
				bins: []Bin{
					&bin{center: "10.000"},
					&bin{center: "12.589"},
					&bin{center: "15.849"},
					&bin{center: "19.953"},
					&bin{center: "25.119"},
					&bin{center: "31.623"},
					&bin{center: "39.811"},
					&bin{center: "50.119"},
					&bin{center: "63.096"},
					&bin{center: "79.433"},
					&bin{center: "100.000"},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad Path: even number of bins",
			args: args{
				fmin:    10,
				fmax:    100,
				qtyBins: 10,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBins(tt.args.fmin, tt.args.fmax, tt.args.qtyBins)

			if tt.wantErr {
				assert.Error(t, err, "expected error")
			} else {
				assert.Nil(t, err, "expected no error")
			}

			assert.Equal(t, tt.want, got, "expected equal bins")
		})
	}
}
