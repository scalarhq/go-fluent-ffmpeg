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

	desired := []string{"-f", "avi", "-i", "pipe:0", "-f", "mp4", "pipe:1"}
	args := fluentffmpeg.NewCommand("").
		PipeInput(buff).
		FromFormat("avi").
		OutputFormat("mp4").
		PipeOutput(buff).GetArgs()

	if !reflect.DeepEqual(args, desired) {
		t.Errorf("Got wrong arguments. Expected: \"%s\", but got: \"%s\"", strings.Join(desired, " "), strings.Join(args, " "))
	}
}
