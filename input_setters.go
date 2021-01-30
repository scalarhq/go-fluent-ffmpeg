package fluentffmpeg

import "io"

// InputPath sets the path of the input file
func (c *Command) InputPath(v string) *Command {
	c.Args.input.inputPath = v

	return c
}

// PipeInput sets the input to be read from an io.Reader
func (c *Command) PipeInput(input io.Reader) *Command {
	c.Args.input.pipeInput = input != nil
	c.input = input

	return c
}

// FromFormat sets the format of the input
func (c *Command) FromFormat(format string) *Command {
	c.Args.input.fromFormat = format
	return c
}
