package fluentffmpeg

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/pkg/errors"
)

var ffprobePath = "ffprobe"

// Probe runs ffprobe with the filePath as input and returns the response as JSON
func Probe(filePath string) (map[string]interface{}, error) {
	cmd := exec.Command(ffprobePath, "-of", "json", "-show_streams", "-show_format", filePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "Running ffprobe failed")
	}
	response := make(map[string]interface{})
	err = json.NewDecoder(stdout).Decode(&response)
	if err != nil {
		errors.Wrap(err, "Failed to decode ffprobe data")
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
		return nil, errors.Wrap(err, "Running ffprobe failed")
	}

	return response, nil
}

// SetFfProbePath sets the path for the ffprobe executable
func SetFfProbePath(path string) {
	ffprobePath = path
}
