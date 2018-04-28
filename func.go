package main

type Func struct {
	pkg string
}

var _ Exporter = Func{}

func (Func) Type() string {
	return "func"
}

func (f Func) Package() string {
	return f.pkg
}

func (Func) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}
