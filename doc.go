/*
	1. checkout last tag in vcs
		1. if no tags then provide user with feedback to say publish 1.0.0 or a beta version 0.1.0
			1aa. check if older version exists first then short circuit the whole thing
		2. find root of repo
		3. rip through all files at current level then recurse into any subpackages
	2. create AST from VCS controlled repo at current state
		1. find root of repo
		2. rip through all files at current level then recurse into any subpackages
	3. create AST of older tag
	4. create store of all packages' exported structs/funcs/consts/vars for current version vs older version(when necessary)
		1. put all types into a `packageStore` by package
			1. Package store should store structs by packageName -> structName -> custom Exporter type, where the Exporter
			     type is an interface that provides 3 basic methods

	type Exporter interface {
		// Type returns type of exported item, be it struct/func/const/var
		Type() string
		// Package returns name of package that this Exporter is exported from
		Package() string
		// Compare compares the Exporter with another Exporter to determine if it is the same or not
		Compare(otherExporter Exporter) (Differences, bool)
	}
		2. Exporter types would be as follows:
			1. Var
			2. Const
			3. Func
				1. Func should have all input/output params (with types associated)
			4. Struct
				1. Struct should have all exported fields and methods embedded in it
					1. Methods should follow the same pattern as Func with the addition of association of method on value or pointer
				2. Struct should have full path included to package, ex. $GOPATH/path_to_package, if unavailable then
				   use root as base level and everything nested underneath it
				3. Struct should have all embedded types listed by full path to struct.
			5. Method
				1. Have all the same things as Func with the addition of bool to indicate caller is value or pointer
			6. Interface
				1. Have signatures of all methods with associated input/output params
			7. Named Func Types
				1. Have signatures of all input/output params and associated types
			8. Custom Types
				1. Have signatures of all input/output params and associated types
			9. Alias Types
*/
package main
