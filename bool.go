package serendipity

func (r *Serendipity) Bool(likeliness ...float64) bool {
	l := .5
	if len(likeliness) > 0 {
		l = likeliness[0]
	}
	if l < 0 || 1 < l {
		return false
	}
	return r.Float64() < l
}
