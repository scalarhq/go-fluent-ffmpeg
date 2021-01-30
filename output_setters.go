package fluentffmpeg

import "io"

// AspectRatio gets the aspect ratio.
// Ex: "16:9"
func (c *Command) AspectRatio(v string) *Command {
	c.Args.output.aspectRatio = v

	return c
}

// Resolution gets the resolution of the media. Ex: "100x100"
func (c *Command) Resolution(v string) *Command {
	c.Args.output.resolution = v

	return c
}

// VideoBitRate gets the video bit rate.
func (c *Command) VideoBitRate(v int) *Command {
	c.Args.output.videoBitRate = v

	return c
}

// VideoMaxBitrate gets the max bit rate for the video.
func (c *Command) VideoMaxBitrate(v int) *Command {
	c.Args.output.videoMaxBitrate = v

	return c
}

// VideoMinBitrate gets the max bit rate for the video.
func (c *Command) VideoMinBitrate(v int) *Command {
	c.Args.output.videoMinBitrate = v

	return c
}

// VideoBitRateTolerance gets the video bit rate tolerance.
func (c *Command) VideoBitRateTolerance(v int) *Command {
	c.Args.output.videoBitRateTolerance = v

	return c
}

// VideoCodec gets the desired video codec when working with video.
func (c *Command) VideoCodec(v string) *Command {
	c.Args.output.videoCodec = v

	return c
}

// VFrames sets the number of frames to output
func (c *Command) VFrames(v int) *Command {
	c.Args.output.vFrames = v

	return c
}

// FrameRate sets the frames per second
func (c *Command) FrameRate(v int) *Command {
	c.Args.output.frameRate = v

	return c
}

// AudioRate sets the audio sampling rate
func (c *Command) AudioRate(v int) *Command {
	c.Args.output.audioRate = v

	return c
}

// AudioBitRate sets the audio bit rate
func (c *Command) AudioBitRate(v int) *Command {
	c.Args.output.audioRate = v

	return c
}

// NativeFramerateInput sets the native frame rate
func (c *Command) NativeFramerateInput(v bool) *Command {
	c.Args.input.nativeFramerateInput = v

	return c
}

// Preset sets the preset
func (c *Command) Preset(v string) *Command {
	c.Args.output.preset = v

	return c
}

// BufferSize sets the buffer size
func (c *Command) BufferSize(v int) *Command {
	c.Args.output.bufferSize = v

	return c
}

// PixelFormat sets the pixel format.
func (c *Command) PixelFormat(v string) *Command {
	c.Args.output.pixelFormat = v

	return c
}

// KeyframeInterval sets the keyframe interval.
func (c *Command) KeyframeInterval(v int) *Command {
	c.Args.output.keyframeInterval = v

	return c
}

// AudioCodec sets the audio codec to use
func (c *Command) AudioCodec(v string) *Command {
	c.Args.output.audioCodec = v

	return c
}

// AudioChannels sets the number of audio channels to use.
func (c *Command) AudioChannels(v int) *Command {
	c.Args.output.audioChannels = v

	return c
}

// OutputFormat sets the format of the output
func (c *Command) OutputFormat(v string) *Command {
	c.Args.output.format = v

	return c
}

// Quality sets the quality
func (c *Command) Quality(v int) *Command {
	c.Args.output.quality = v

	return c
}

// OutputPath sets the path to write the output file
func (c *Command) OutputPath(v string) *Command {
	c.Args.output.outputPath = v

	return c
}

// PipeOutput sets the output to be written to an io.Writer
func (c *Command) PipeOutput(output io.Writer) *Command {
	c.Args.output.pipeOutput = output != nil
	c.output = output

	return c
}

// Overwrite configures FFmpeg to overwrite existing files
func (c *Command) Overwrite(b bool) *Command {
	c.Args.output.overwrite = b

	return c
}

// OutputOptions sets additional output options
func (c *Command) OutputOptions(options ...string) *Command {
	c.Args.output.outputOptions = options

	return c
}
