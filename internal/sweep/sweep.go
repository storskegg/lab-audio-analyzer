package sweep

import (
	"context"
	"fmt"
	"sync"
)

type Incrementation int

const (
	IncrementationDefault Incrementation = iota
	IncrementationLogarithmic
	IncrementationLinear
)

type Config struct {
	Buckets        int
	Incrementation Incrementation
	FrequencyStart float64
	FrequencyEnd   float64
}

type Sweep interface {
	Once() error
	Start() error
	Stop() error
	RemoveBucket(idx int)
	Measure(freq float64) (float64, error)
}

func NewSweep(ctx context.Context, config *Config) Sweep {
	s := sweep{
		Config:  config,
		Buckets: make(map[int]Bucket),
	}

	s.ctx, s.cancel = context.WithCancel(ctx)

	bucketWidth := (s.Config.FrequencyEnd - s.Config.FrequencyStart) / float64(s.Config.Buckets)

	for i := 0; i < s.Config.Buckets; i++ {
		start := s.Config.FrequencyStart + float64(i)*bucketWidth
		end := start + bucketWidth

		s.Buckets[i] = NewBucket(start, end)
	}

	return &s
}

type sweep struct {
	ctx    context.Context
	cancel context.CancelFunc

	muBuckets sync.Mutex
	muMeasure sync.Mutex

	Config  *Config
	Buckets map[int]Bucket
}

func (s *sweep) Once() error {
	return fmt.Errorf("not implemented yet")
}

func (s *sweep) Start() error {
	for idx, b := range s.Buckets {
		if b == nil {
			// if we run into a nil bucket in the map, then remove it and start
			// over; this will not infinite loop unless you have infinite
			// buckets, and ja...fat chance of that...
			s.RemoveBucket(idx)
			return s.Start()
		}

		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		default:
			s.muMeasure.Lock()

			vrms, err := s.Measure(b.Center())
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				b.UpdateVrms(vrms)
			}

			s.muMeasure.Unlock()
		}
	}

	return nil
}

func (s *sweep) Stop() error {
	s.cancel()

	return s.ctx.Err()
}

func (s *sweep) RemoveBucket(idx int) {
	s.muBuckets.Lock()
	defer s.muBuckets.Unlock()

	delete(s.Buckets, idx)
}

func (s *sweep) Measure(freq float64) (float64, error) {
	// return device.Measure(freq)

	return -1, nil
}

type Bucket interface {
	Start() float64
	Center() float64
	End() float64
	Vrms() float64

	UpdateStart(start float64)
	UpdateEnd(end float64)
	UpdateVrms(vRms float64)
}

func NewBucket(start, end float64) Bucket {
	return &bucket{
		start: start,
		end:   end,
	}
}

type bucket struct {
	mu sync.RWMutex

	start, end float64
	vRms       float64
}

func (b *bucket) Start() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.start
}

func (b *bucket) Center() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return (b.start + b.end) / 2
}

func (b *bucket) End() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.end
}

func (b *bucket) Vrms() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.vRms
}

func (b *bucket) UpdateStart(start float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.start = start
}

func (b *bucket) UpdateEnd(end float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.end = end
}

func (b *bucket) UpdateVrms(vRms float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.vRms = vRms
}
