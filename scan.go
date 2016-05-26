package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"sourcegraph.com/sourcegraph/srclib/unit"
)

type ScanCmd struct{}

//An external configuration file is not needed for JSON, and so
//all logic relating to it is not included in this file.

var (
	parser  = flags.NewNamedParser("srclib-json", flags.Default)
	scanCmd = ScanCmd{}
)

func init() {
	_, err := parser.AddCommand("scan",
		"scan for JSON files",
		"Scan the directory tree rooted at the current directory for JSON Files",
		&scanCmd)

	if err != nil {
		log.Fatal(err)
	}

}

// func main() {
// 	if _, err := parser.Parse(); err != nil {
// 		os.Exit(1)
// 	}
// }

func isJSONFile(fileName string) bool {
	return filepath.Ext(fileName) == ".json"
}

func (c *ScanCmd) Execute(args []string) error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	units, err := scan(cwd)

	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(units, "", " ")

	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(out)

	if err != nil {
		return err
	}

	return nil

}

func scan(dir string) ([]*unit.SourceUnit, error) {
	u := unit.SourceUnit{
		Name: filepath.Base(dir),
		Type: "json",
	}
	units := []*unit.SourceUnit{&u}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if isJSONFile(path) {
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			u.Files = append(u.Files, filepath.ToSlash(relPath))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return units, nil
}
