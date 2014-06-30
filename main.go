package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage:", args[0], "[path/to/latencies.csv]")
		os.Exit(1)
	}

	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	distribution := &Distribution{}

	// Start from 1 to skip header row.
	for _, row := range rows[1:] {
		// Lower Bound, Upper Bound (non-inclusive), Count, Probability.
		var floats [4]float64

		for idx := range row {
			float, err := strconv.ParseFloat(row[idx], 64)
			if err != nil {
				log.Fatal(err)
			}
			floats[idx] = float
		}

		distribution.Add(floats[0], floats[1], floats[2], floats[3])
	}
}
