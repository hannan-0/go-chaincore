package cap

// VectorClock represents a simplified vector clock
type VectorClock map[string]int

// Update updates the vector clock with a new event
func (vc VectorClock) Update(nodeID string) {
	vc[nodeID]++
}

// Compare returns:
// -1 if vc < other (happened-before),
//
//	1 if vc > other (happened-after),
//	0 if vc and other are concurrent
func (vc VectorClock) Compare(other VectorClock) int {
	selfLess := false
	otherLess := false

	allKeys := make(map[string]struct{})
	for k := range vc {
		allKeys[k] = struct{}{}
	}
	for k := range other {
		allKeys[k] = struct{}{}
	}

	for k := range allKeys {
		v1 := vc[k]
		v2 := other[k]
		if v1 < v2 {
			selfLess = true
		} else if v1 > v2 {
			otherLess = true
		}
	}

	if selfLess && !otherLess {
		return -1
	} else if otherLess && !selfLess {
		return 1
	}
	return 0
}

// DetectConflict returns true if the clocks are concurrent
func DetectConflict(vc1, vc2 VectorClock) bool {
	return vc1.Compare(vc2) == 0
}

// ResolveConflict uses entropy (string length here) as tiebreaker
func ResolveConflict(data1, data2 string) string {
	if len(data1) >= len(data2) {
		return data1
	}
	return data2
}
