package core

import "fmt"

type Parameters struct {
	Loudness string
	TruePeak string
	Range    string
}

func (parameters Parameters) String() string {
	return fmt.Sprintf("I=%s:TP=%s:LRA=%s", parameters.Loudness, parameters.TruePeak, parameters.Range)
}

type Analysis struct {
	InputLoudness  string `json:"input_i"`
	InputTruePeak  string `json:"input_tp"`
	InputRange     string `json:"input_lra"`
	InputThreshold string `json:"input_thresh"`
}

func (analysis Analysis) String() string {
	return fmt.Sprintf("measured_I=%s:measured_TP=%s:measured_LRA=%s:measured_thresh=%s", analysis.InputLoudness, analysis.InputTruePeak, analysis.InputRange, analysis.InputThreshold)
}
