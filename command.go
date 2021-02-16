package fluentffmpeg

import (
	"context"
	"io"
	"os/exec"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

// Command is a struct that holds arguments and their values to run FFmpeg
type Command struct {
	FFmpegPath string
	Args       *Args
	input      io.Reader
	output     io.Writer
	logs       io.Writer
}

// NewCommand returns a new Command
func NewCommand(ffmpegPath string) *Command {
	if ffmpegPath == "" {
		ffmpegPath = "ffmpeg"
	}
	return &Command{
		FFmpegPath: ffmpegPath,
		Args: &Args{
			output: outputArgs{
				constantRateFactor: -1, // Initialize to -1 because zero value is a valid parameter
			},
		},
	}
}

// Run runs the FFmpeg command. It returns an error if the command fails with exit status code 1. This error message only signifies that
// the command returned a non-zero status code, read from stderr to see more comprehensive FFmpeg errors.
func (c *Command) Run() error {
	return c.Build().Run()
}

// RunWithContext is like Run but includes a context which is used to kill the process
func (c *Command) RunWithContext(ctx context.Context) error {
	return c.BuildWithContext(ctx).Run()
}

// Build returns an exec.Cmd struct ready to run the FFmpeg command with its arguments
func (c *Command) Build() *exec.Cmd {
	return c.BuildWithContext(context.Background())
}

// BuildWithContext is like Build but includes a context which is used to kill the process
func (c *Command) BuildWithContext(ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.FFmpegPath, c.GetArgs()...)

	if c.input != nil {
		cmd.Stdin = c.input
	}

	if c.output != nil {
		cmd.Stdout = c.output
	}

	if c.logs != nil {
		cmd.Stderr = c.logs
	}

	return cmd
}

// GetArgs returns the arguments for the FFmpeg command.
func (c *Command) GetArgs() []string {
	var options []string

	options = append(options, c.getArgs(c.Args.input, "pipeInput", "inputPath")...)
	options = append(options, c.getArgs(c.Args.output, "pipeOutput", "outputPath")...)

	return append(options, c.Args.globalOptions...)
}

func (c *Command) getArgs(argType interface{}, targetNames ...string) []string {
	var options []string
	var target []string

	fields := structs.Names(argType)

	// Iterates through the fields,
	// and calls its corresponding getter function.
	for _, v := range fields {
		option := true
		if containsString(targetNames, v) {
			option = false
		}
		value := reflect.ValueOf(c.Args).MethodByName("Get" + strings.Title(v))
		if (value != reflect.Value{}) {
			result := value.Call([]reflect.Value{})
			if v, ok := result[0].Interface().([]string); ok {
				if option {
					options = append(options, v...)
				} else {
					target = append(target, v...)
				}
			}
		}
	}
	return append(options, target...)
}

// OutputLogs sets the destination to write the FFmpeg log output to
func (c *Command) OutputLogs(writer io.Writer) *Command {
	c.logs = writer
	return c
}
