package main

type NamedFunc struct {
	pkg string
}

var _ Exporter = NamedFunc{}

func (NamedFunc) Type() string {
	return "namedFunc"
}

func (n NamedFunc) Package() string {
	return n.pkg
}

func (NamedFunc) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}
