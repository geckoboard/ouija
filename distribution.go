package main

type Distribution struct {
	ranges []Range
}

type Range struct {
	From  float64
	To    float64
	Count float64
	Prob  float64
}

func (r *Range) includes(point float64) bool {
	return (point > r.From && point <= r.To)
}

func (d *Distribution) Add(from, to, count, prob float64) {
	d.ranges = append(d.ranges, Range{From: from, To: to, Count: count, Prob: prob})
}

func (d Distribution) Output(base float64) float64 {
	var start, end float64

	for _, r := range d.ranges {
		end += r.Prob

		if base > start && base < end {
			return r.From
		}

		start = end
	}

	return 0.0
}
