package main

import (
	"audio-normalize/cmd"
	"audio-normalize/internal/core"
	"audio-normalize/pkg/logs"
)

func main() {
	input := cmd.GetInput()
	parameters := core.Parameters{
		Loudness: cmd.GetI(),
		TruePeak: cmd.GetTP(),
		Range:    cmd.GetLRA(),
	}

	logs.Info("main", "analyzing the input file")

	analysis, err := core.Analyze(input, parameters)
	if err != nil {
		logs.Fatal("core.Analyze", "couldn't analyze the input file", err)
		return
	}

	logs.Info("main", "normalizing the input file")

	err = core.Normalize(input, parameters, analysis)
	if err != nil {
		logs.Fatal("core.Normalize", "couldn't normalize the input file", err)
		return
	}
}
