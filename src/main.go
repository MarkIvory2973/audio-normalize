package main

import (
	"audio-normalize/cmd"
	"audio-normalize/internal/core"
	"audio-normalize/pkg/logs"
)

func main() {
	inputs, err := cmd.GetInputs()
	if err != nil {
		logs.Fatal("cmd.GetInputs", "couldn't get inputs", err)
		return
	}

	parameters := core.Parameters{
		Loudness: cmd.GetI(),
		TruePeak: cmd.GetTP(),
		Range:    cmd.GetLRA(),
	}

	for _, input := range inputs {
		logs.Info("main", "analyzing %q", input)

		analysis, err := core.Analyze(input, parameters)
		if err != nil {
			logs.Fatal("core.Analyze", "couldn't analyze %q", err, input)
			return
		}

		logs.Info("main", "normalizing %q", input)

		err = core.Normalize(input, parameters, analysis)
		if err != nil {
			logs.Fatal("core.Normalize", "couldn't normalize %q", err, input)
			return
		}
	}
}
