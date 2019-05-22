package set

func (d *datum) Union(d2 *datum) *datum {
	return Union(d, d2)
}

func (d *datum) Intersection(d2 *datum) *datum {
	return Intersection(d, d2)
}

func (d *datum) Difference(d2 *datum) *datum {
	return Difference(d, d2)
}

func (d *datum) Subset(d2 *datum) bool {
	return Subset(d, d2)
}
