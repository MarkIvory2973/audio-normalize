package core

import (
	"audio-normalize/pkg/jsons"
	"bytes"
	"fmt"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func Analyze(input string, parameters Parameters) (Analysis, error) {
	filter := fmt.Sprintf("loudnorm=%s:print_format=json", parameters.String())
	buffer := bytes.Buffer{}
	err := ffmpeg_go.
		Input(input, ffmpeg_go.KwArgs{
			"hide_banner": "",
		}).
		Output("-", ffmpeg_go.KwArgs{
			"af": filter,
			"f":  "null",
		}).
		WithErrorOutput(&buffer).
		Silent(true).
		Run()
	if err != nil {
		return Analysis{}, err
	}

	var analysis Analysis
	err = jsons.Unmarshal(buffer.Bytes(), &analysis)
	if err != nil {
		return Analysis{}, err
	}

	return analysis, nil
}
