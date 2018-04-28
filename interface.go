package main

type Interface struct {
	pkg string
}

var _ Exporter = Interface{}

func (Interface) Type() string {
	return "interface"
}

func (i Interface) Package() string {
	return i.pkg
}

func (Interface) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}
