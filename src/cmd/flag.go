package cmd

import (
	"flag"
	"path/filepath"
)

var inputs string
var i string
var tp string
var lra string

func init() {
	flag.StringVar(&inputs, "inputs", "", "Specify the input files")
	flag.StringVar(&i, "i", "-24", "Specify the loudness")
	flag.StringVar(&tp, "tp", "-2", "Specify the true peak")
	flag.StringVar(&lra, "lra", "25", "Specify the range")
	flag.Parse()
}

func GetInputs() ([]string, error) {
	inputs, err := filepath.Glob(inputs)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

func GetI() string {
	return i
}

func GetTP() string {
	return tp
}

func GetLRA() string {
	return lra
}
