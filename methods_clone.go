package set

func (d *datum) Clone() *datum {
	return Clone(d)
}
