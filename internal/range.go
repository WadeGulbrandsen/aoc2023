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

func (r Range) Len() int {
	return AbsDiff(r.Start, r.End) + 1
}

func (r Range) Contains(i int) bool {
	return i >= r.Start && i <= r.End
}

func (r Range) Intersect(other Range) bool {
	return r.Contains(other.Start) || r.Contains(other.End)
}

func (r Range) Adjacent(other Range) bool {
	return r.End+1 == other.Start || other.End+1 == r.Start
}

func (r Range) Combine(other Range) (Range, bool) {
	var combined Range
	if r.Intersect(other) || r.Adjacent(other) {
		combined.Start = min(r.Start, other.Start)
		combined.End = max(r.Start, other.End)
		return combined, true
	}
	return combined, false
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

func (r Range) SplitAt(i int) (Range, Range) {
	var lower, higher Range
	switch {
	case i <= r.Start:
		higher = r
	case i > r.End:
		lower = r
	default:
		lower = Range{r.Start, i - 1}
		higher = Range{i, r.End}
	}
	return lower, higher
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

func (r *RangeList) Compact() RangeList {
	n := slices.Clone(*r)
	n.Sort()
	if n.Len() < 2 {
		return n
	}
	compacted := RangeList{n[0]}
	for _, o := range n[1:] {
		i := compacted.Len() - 1
		if combined, e := compacted[i].Combine(o); e {
			compacted[i] = combined
		} else {
			compacted = append(compacted, o)
		}
	}
	return compacted
}
