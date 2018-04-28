package main

type Exporter interface {
	// Type returns the type , be it var, const, func, struct, or method
	Type() string
	// Package returns the parent package of the type
	Package() string
	// Compare compares the type to another Exporter type to determine if the
	// two types match completely. Any differences are collected in the Differences
	// return type and the match flag will be set to false
	Compare(Exporter) (diffs Differences, match bool)
}

type Differences struct {
	D []Difference
}

type Difference struct {
	oldLine, newLine string
}