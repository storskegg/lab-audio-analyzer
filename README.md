# lab-audio-analyzer

Lab tool; measure audio response of an audio device (filter, preamp, etc).

## A Sweep

- Start Frequency
- End Frequency
- Number of Buckets

### Process -- `foreach Bucket b`

1. `device.DDS.SetFrequency(b.Center())`
2. `wait until device.CountFrequency() == b.Center()`
3. `device.MeasureVrms()`

## Bucket

- Start Frequency
- End Frequency
- `Center()` Frequency

