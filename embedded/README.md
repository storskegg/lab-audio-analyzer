# `embedded`: A Pivot

Originally, I was going to write a Go application that would control my oscilloscope:

- Set a DDS frequency at X Vpp
- Measure Vrms, and report
- Rinse, repeat.

This would have been a horribly slow process in order to, say, sweep 20 Hz to 20 kHz, and
measuring 4001 bins. Additionally, it would have reported Vrms total, and not Vrms of the
fundamental, and also its harmonics (and there will be harmonics).

I need to do real FFT work, and I need to do this as firmware for a device to reduce
latencies, and increase control and visibility.

## Hardware Decisions

- Arduino GIGA R1 M7
  - **What:** One hell of a board boasting one hell of an mcu: an STM32H7xx beast
    of a controller.
  - **Why:** This thing has more than enough floating point capability for running the FFTs
    I'll be using. It also has strong graphics capabilities, support for the display shield,
    an onboard, 16-bit ADC that isn't complete trash, and that is already wired up to the
    microphone ring of the onboard 1/8" audio TRRS jack. (I can add better DAQ later.)
- Arduino GIGA Display
  - **What:** It's all in the name
  - **Why:** High (enough) resolution display with touch support, and ready-made for the host
    controller (complete with graphics acceleration, etc)
- AD9833
  - **What:** Analog Devices DDS waveform generator.
  - **Why:** Ja, so the STM32H7xx mcu has a decent-enough DAC onboard that could be used for
    sweep/marker generation, but I wanted a little more resolution, and more importantly I
    wanted to defer the actual DDS function to dedicated hardware vs occupying precious
    cpu cycles to math/output sinusoidal waveforms.