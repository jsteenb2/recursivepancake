package main

// Alias represents an alias type in go. Added in go v1.9 these types
// simply provide the ability to alias an existing type in go. Declerations like
// this are like so:
// type T1 = T2
type Alias struct {
	pkg string
}

var _ Exporter = Alias{}

func (Alias) Type() string {
	return "alias"
}

func (a Alias) Package() string {
	return a.pkg
}

func (Alias) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}
