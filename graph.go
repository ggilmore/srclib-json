package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"sourcegraph.com/sourcegraph/srclib-json/myjson"
)

const testJSON = `{"name": "Geoffrey", "parents": ["Donna", "Henry"], "age": 22, "clothes": {"shirt": true, "pants": "true"} }`

func main() {
	dec := myjson.NewDecoder(strings.NewReader(testJSON))
	for {
		t, start, stop, path, err := dec.EndpToken()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%T: %v, start: %d, stop %d, path: %v", t, t, start, stop, path)
		fmt.Printf("\n")

	}
}
