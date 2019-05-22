package set

func (d *Datum) Clone() *Datum {
	return Clone(d)
}
