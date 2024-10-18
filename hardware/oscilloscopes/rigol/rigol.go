// rigol package provides abstration for rigol oscilloscopes. this is the first
// supported series of devices, because this is what I have in my lab.

package rigol

import "github.com/storskegg/lab-audio-analyzer/hardware/functions"

const Functions = functions.SupportsFunctionVrms & functions.SupportsFunctionDDS
