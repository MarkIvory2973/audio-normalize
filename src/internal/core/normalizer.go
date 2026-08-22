package core

import (
	"fmt"
	"os"
	"path/filepath"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func Normalize(input string, parameters Parameters, analysis Analysis) error {
	err := os.MkdirAll("dist", os.ModePerm)
	if err != nil {
		return err
	}

	_, name := filepath.Split(input)
	output := filepath.Join("dist", name)
	filter := fmt.Sprintf("loudnorm=%s:%s:linear=true", parameters.String(), analysis.String())
	err = ffmpeg_go.
		Input(input, ffmpeg_go.KwArgs{
			"hide_banner": "",
		}).
		Output(output, ffmpeg_go.KwArgs{
			"af": filter,
		}).
		Silent(true).
		Run()
	if err != nil {
		return err
	}

	return nil
}
