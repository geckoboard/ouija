package main

import (
	"testing"
)

func TestDistributionAdd(t *testing.T) {
	dist := Distribution{}
	dist.Add(0, 5, 25, 0.25)
	dist.Add(5, 10, 75, 0.75)

	number := dist.Output(0.9)

	if number >= 10 || number < 5 {
		t.Fatal("distribution returned a number in the wrong range")
	}
}
