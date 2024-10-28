package binning

import "sync"

type Bin interface {
	Start() float64
	End() float64
	Width() float64

	UpdateBounds(start, end float64)
}

func NewBin(start, end float64) Bin {
	return &bin{
		start: start,
		end:   end,
	}
}

type bin struct {
	mu sync.RWMutex

	start float64
	end   float64
}

func (b *bin) Start() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.start
}

func (b *bin) End() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.end
}

func (b *bin) Width() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.end - b.start
}

func (b *bin) UpdateBounds(start, end float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.start = start
	b.end = end
}
