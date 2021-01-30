package fluentffmpeg_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func TestArgsOrder(t *testing.T) {
	buff := &bytes.Buffer{}

	desired := []string{"-f", "avi", "-test", "123", "-i", "pipe:0", "-f", "mp4", "-movflags", "empty_moov", "pipe:1", "-report"}
	args := fluentffmpeg.NewCommand("").
		PipeInput(buff).
		InputOptions("-test", "123").
		FromFormat("avi").
		OutputFormat("mp4").
		OutputOptions("-movflags", "empty_moov").
		Options("-report").
		PipeOutput(buff).GetArgs()

	if !reflect.DeepEqual(args, desired) {
		t.Errorf("Got wrong arguments. Expected: \"%s\", but got: \"%s\"", strings.Join(desired, " "), strings.Join(args, " "))
	}
}
