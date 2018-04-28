package main

type Const struct {
	value interface{}
	pkg string
}

var _ Exporter = Const{}

func (Const) Type() string {
	return "const"
}

func (c Const) Package() string {
	return c.pkg
}

func (Const) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}