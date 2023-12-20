package interval

import (
	"AOC/lib"
	"fmt"
	"math"
)

type Interval struct {
	Start int64
	End   int64
}

func (i *Interval) IsFullyContained(other Interval) bool {
	return i.Start >= other.Start && i.End <= other.End
}

func (i *Interval) IsOverlap(other Interval) bool {
	return other.Start <= i.End && i.Start <= other.End
}

func (i *Interval) Overlap(other Interval) (Interval, bool) {
	if i.IsOverlap(other) {
		return Interval{Start: lib.Max(i.Start, other.Start), End: lib.Min(i.End, other.End)}, true
	}

	return Interval{Start: math.MinInt64, End: math.MinInt64}, false
}

func (i *Interval) Split(point int64) []Interval {
	if point < i.Start || point > i.End {
		return []Interval{{Start: i.Start, End: i.End}}
	}

	return []Interval{{Start: i.Start, End: point - 1}, {Start: point, End: i.End}}
}

func (i *Interval) Clone() *Interval {
	return &Interval{Start: i.Start, End: i.End}
}

func (i *Interval) Len() int64 {
	if i.End > i.Start {
		return i.End - i.Start + 1
	} else if i.Start > i.End {
		return i.Start - i.End + 1
	}
	return 0
}

func (i *Interval) Compare(other Interval) int {
	if i.Start < other.Start {
		return -1
	}
	if i.Start > other.Start {
		return 1
	}
	if i.End < other.End {
		return -1
	}
	if i.End > other.End {
		return 1
	}
	return 0
}

func (i *Interval) String() string {
	return fmt.Sprintf("%d..%d", i.Start, i.End)
}
