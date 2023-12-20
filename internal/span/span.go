package span

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

type Span struct {
	Start, End int
}

func CompareRanges(a, b Span) int {
	if n := cmp.Compare(a.Start, b.Start); n != 0 {
		return n
	}
	return cmp.Compare(a.End, b.End)
}

func (r Span) String() string {
	return fmt.Sprintf("[%v-%v]", r.Start, r.End)
}

func (r Span) IsEmpty() bool {
	return r.Start == 0 && r.End == 0
}

func (r Span) Len() int {
	return utils.AbsDiff(r.Start, r.End) + 1
}

func (r Span) Contains(i int) bool {
	return i >= r.Start && i <= r.End
}

func (r Span) Intersect(other Span) bool {
	return r.Contains(other.Start) || r.Contains(other.End)
}

func (r Span) Adjacent(other Span) bool {
	return r.End+1 == other.Start || other.End+1 == r.Start
}

func (r Span) Combine(other Span) (Span, bool) {
	var combined Span
	if r.Intersect(other) || r.Adjacent(other) {
		combined.Start = min(r.Start, other.Start)
		combined.End = max(r.Start, other.End)
		return combined, true
	}
	return combined, false
}

func (r Span) SplitOtherRange(other Span) (Span, Span, Span) {
	var before, contained, after Span
	switch {
	case other.End < r.Start:
		before = other
	case other.Start > r.End:
		after = other
	case other.Start >= r.Start && other.End <= r.End:
		contained = other
	case other.Start < r.Start && other.End <= r.End:
		before = Span{other.Start, r.Start - 1}
		contained = Span{r.Start, other.End}
	case other.Start >= r.Start && other.End > r.End:
		contained = Span{other.Start, r.End}
		after = Span{r.End + 1, other.End}
	default:
		before = Span{other.Start, r.Start - 1}
		contained = r
		after = Span{r.End + 1, other.End}
	}
	return before, contained, after
}

func (r Span) SplitAt(i int) (Span, Span) {
	var lower, higher Span
	switch {
	case i <= r.Start:
		higher = r
	case i > r.End:
		lower = r
	default:
		lower = Span{r.Start, i - 1}
		higher = Span{i, r.End}
	}
	return lower, higher
}

type SpanList []Span

func (r SpanList) Len() int {
	return len(r)
}

func (r SpanList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r SpanList) Less(i, j int) bool {
	return CompareRanges(r[i], r[j]) < 0
}

func (r SpanList) Sort() {
	slices.SortFunc(r, CompareRanges)
}

func (r SpanList) IsSorted() bool {
	return slices.IsSortedFunc(r, CompareRanges)
}

func (r SpanList) FilterEmpty() SpanList {
	return slices.DeleteFunc(r, func(r Span) bool { return r.IsEmpty() })
}

func (r *SpanList) Compact() SpanList {
	n := slices.Clone(*r)
	n.Sort()
	if n.Len() < 2 {
		return n
	}
	compacted := SpanList{n[0]}
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
