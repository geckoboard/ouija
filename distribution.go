package main

import (
	"math/rand"
	"time"
)

type Distribution struct {
	ranges []Range
	random *rand.Rand
}

type Range struct {
	From  float64
	To    float64
	Count float64
	Prob  float64
}

func NewDistribution() *Distribution {
	randSrc := rand.NewSource(time.Now().UTC().UnixNano())
	return &Distribution{
		random: rand.New(randSrc),
	}
}

func (d *Distribution) Add(from, to, count, prob float64) {
	rng := Range{From: from, To: to, Count: count, Prob: prob}
	d.ranges = append(d.ranges, rng)
}

func (d Distribution) Output(base float64) float64 {
	var start, end float64

	for _, rng := range d.ranges {
		end += rng.Prob

		if base > start && base < end {
			return d.randomBetween(rng.From, rng.To)
		}

		start = end
	}

	return 0.0
}

// Assumes that start <= end.
func (d Distribution) randomBetween(start, end float64) float64 {
	n := d.random.Float64()
	return n * (end - start) + start
}
