package main

import (
	"testing"
)

func TestDistributionAdd(t *testing.T) {
	dist := NewDistribution()
	dist.Add(0, 5, 25, 0.25)
	dist.Add(5, 10, 75, 0.75)

	// rand number ideally
	number := dist.Output(0.9)

	if number >= 10 || number < 5 {
		t.Fatalf("distribution returned a number in the wrong range: %v", number)
	}

}

func TestDistributionRandomness(t *testing.T) {
	dist := NewDistribution()
	dist.Add(0, 10, 100, 1.0)

	if dist.Output(0.5) == dist.Output(0.5) {
		t.Fatal("distribution returned the same number in a row")
	}
}
