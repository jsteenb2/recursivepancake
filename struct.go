package main

type Struct struct {
	pkg string
}

var _ Exporter = Struct{}

func (Struct) Type() string {
	return "struct"
}

func (s Struct) Package() string {
	return s.pkg
}

func (Struct) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}

type Method struct {
	pkg string
	pointerCaller bool
}

var _ Exporter = Method{}

func (Method) Type() string {
	return "method"
}

func (m Method) Package() string {
	return m.pkg
}

func (Method) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}