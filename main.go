package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml2json "gopkg.in/peter-edge/yaml2json.v1"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("Couldn't read input: %v\n", err)
		os.Exit(1)
	}

	output, err := yaml2json.Convert(input)
	if err != nil {
		fmt.Printf("Couldn't convert input: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stdout.Write([]byte(output))
	if err != nil {
		fmt.Printf("Couldn't write output: %v\n", err)
		os.Exit(1)
	}
}
