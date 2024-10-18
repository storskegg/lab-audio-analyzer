// sweep performs an audio sweep on the DUT

package sweep

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.1"

func init() {
	cmdRoot.Flags().IntVarP(&flagSweepCount, "count", "c", -1, "Number of sweeps (Default: continuous)")
	cmdRoot.Flags().Float64VarP(&flagFrequencyStart, "start", "s", 1000.0, "Start frequency (Hz)")
	cmdRoot.Flags().Float64VarP(&flagFrequencyEnd, "end", "e", 10000.0, "End frequency (Hz)")
	cmdRoot.Flags().BoolVarP(&flagLogarithmic, "log", "l", true, "Logarithmic sweep (default: true)")
	cmdRoot.Flags().BoolVarP(&flagLinear, "lin", "L", false, "Linear sweep (default: false)")

	cmdRoot.MarkFlagsMutuallyExclusive("logarithmic", "linear")
}

var (
	flagSweepCount     int
	flagFrequencyStart float64
	flagFrequencyEnd   float64
	flagLogarithmic    bool
	flagLinear         bool
)

var cmdRoot = &cobra.Command{
	Use:     "sweep",
	Short:   "sweep",
	Version: Version,
	RunE:    execRoot,
}

func execRoot(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented yet")
}
