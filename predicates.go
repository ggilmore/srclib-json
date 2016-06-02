package main

import "path/filepath"

//Predicates.go contains predicates that check a given file path
//to see if it is a JSON file that we recognize and support.

func npmPrecicate(path string) bool {
	file := filepath.Base(path)
	return file == "package.json"
}

func typescriptPredicate(path string) bool {
	dir, file := filepath.Split(path)
	return file == "tsconfig.json" || file == "tslint.json" || (filepath.Base(dir) == "typings" && file == "typings.json")
}

func meteorPredicate(path string) bool {
	file := filepath.Base(path)
	return file == "settings.json"
}

func init() {
	for _, p := range []func(string) bool{npmPrecicate, typescriptPredicate, meteorPredicate} {
		filePredicates = append(filePredicates, p)
	}
}
