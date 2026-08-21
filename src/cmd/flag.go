package cmd

import "flag"

var input string
var i string
var tp string
var lra string

func init() {
	flag.StringVar(&input, "input", "", "Specify the input file")
	flag.StringVar(&i, "i", "-24", "Specify the loudness")
	flag.StringVar(&tp, "tp", "-2", "Specify the true peak")
	flag.StringVar(&lra, "lra", "25", "Specify the range")
	flag.Parse()
}

func GetInput() string {
	return input
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
