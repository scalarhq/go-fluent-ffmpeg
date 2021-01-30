package fluentffmpeg

// GetInputPath returns the input file path
func (a *Args) GetInputPath() []string {
	if a.input.inputPath != "" {
		if a.input.nativeFramerateInput {
			return []string{"-re", "-i", a.input.inputPath}
		}
		return []string{"-i", a.input.inputPath}
	}
	return nil
}

// GetPipeInput returns whether or not ffmpeg is set to receive piped input
func (a *Args) GetPipeInput() []string {
	if a.input.pipeInput != false {
		return []string{"-i", "pipe:0"}
	}

	return nil
}

// GetFromFormat returns the input format
func (a *Args) GetFromFormat() []string {
	if a.input.fromFormat != "" {
		return []string{"-f", a.input.fromFormat}
	}

	return nil
}
