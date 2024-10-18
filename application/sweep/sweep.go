// sweep performs an audio sweep on the DUT

package sweep

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.1"

var cmdRoot = &cobra.Command{
	Use:     "sweep",
	Short:   "sweep",
	Version: Version,
	RunE:    execRoot,
}

func execRoot(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented yet")
}
