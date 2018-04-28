package main

// CustomType implements a struct that conforms to the Exporter interface. An
// example of a custom type can be something like the following
// type MyCustomType []string or similar
// Or can even be a type of a struct that the api in question the api implements
type CustomType struct {
	pkg string
}

var _ Exporter = CustomType{}

func (CustomType) Type() string {
	return "custom_type"
}

func (c CustomType) Package() string {
	return c.pkg
}

func (CustomType) Compare(Exporter) (Differences, bool) {
	return Differences{}, false
}
