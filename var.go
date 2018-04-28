package main

type Var struct {
	value interface{}
	pkg string
}

var _ Exporter = Var{}

func (Var) Type() string {
	return  "var"
}

func (v Var) Package() string {
	return v.pkg
}

func (Var) Compare(e Exporter) (Differences, bool) {
	return Differences{}, false
}