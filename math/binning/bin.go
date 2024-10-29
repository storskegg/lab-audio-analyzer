package binning

import (
	"fmt"
	"sync"
)

//type Bin interface {
//	Start() float64
//	End() float64
//	Width() float64
//
//	UpdateBounds(start, end float64)
//}

type Bin interface {
	Center() string

	UpdateCenter(center float64)
}

func NewBin(center float64) Bin {
	return &bin{
		center: fmt.Sprintf("%.3f", center),
	}
}

type bin struct {
	mu sync.RWMutex

	center string
	//value  string
}

func (b *bin) Center() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.center
}

func (b *bin) UpdateCenter(center float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.center = fmt.Sprintf("%.3f", center)
}

//func (b *bin) Start() float64 {
//	b.mu.RLock()
//	defer b.mu.RUnlock()
//
//	return b.start
//}
//
//func (b *bin) End() float64 {
//	b.mu.RLock()
//	defer b.mu.RUnlock()
//
//	return b.end
//}
//
//func (b *bin) Width() float64 {
//	b.mu.RLock()
//	defer b.mu.RUnlock()
//
//	return b.end - b.start
//}
//
//func (b *bin) UpdateBounds(start, end float64) {
//	b.mu.Lock()
//	defer b.mu.Unlock()
//
//	b.start = start
//	b.end = end
//}
