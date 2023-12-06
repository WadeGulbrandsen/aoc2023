package internal

import (
	"cmp"
	"fmt"
	"slices"
)

type Range struct {
	Start, End int
}

func CompareRanges(a, b Range) int {
	if n := cmp.Compare(a.Start, b.Start); n != 0 {
		return n
	}
	return cmp.Compare(a.End, b.End)
}

func (r Range) String() string {
	return fmt.Sprintf("[%v-%v]", r.Start, r.End)
}

func (r Range) IsEmpty() bool {
	return r.Start == 0 && r.End == 0
}

func (r Range) SplitOtherRange(other Range) (Range, Range, Range) {
	var before, contained, after Range
	switch {
	case other.End < r.Start:
		before = other
	case other.Start > r.End:
		after = other
	case other.Start >= r.Start && other.End <= r.End:
		contained = other
	case other.Start < r.Start && other.End <= r.End:
		before = Range{other.Start, r.Start - 1}
		contained = Range{r.Start, other.End}
	case other.Start >= r.Start && other.End > r.End:
		contained = Range{other.Start, r.End}
		after = Range{r.End + 1, other.End}
	default:
		before = Range{other.Start, r.Start - 1}
		contained = r
		after = Range{r.End + 1, other.End}
	}
	return before, contained, after
}

type RangeList []Range

func (r RangeList) Len() int {
	return len(r)
}

func (r RangeList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RangeList) Less(i, j int) bool {
	return CompareRanges(r[i], r[j]) < 0
}

func (r RangeList) Sort() {
	slices.SortFunc(r, CompareRanges)
}

func (r RangeList) IsSorted() bool {
	return slices.IsSortedFunc(r, CompareRanges)
}

func (r RangeList) FilterEmpty() RangeList {
	return slices.DeleteFunc(r, func(r Range) bool { return r.IsEmpty() })
}
