package fluentffmpeg

import (
	"fmt"
	"strconv"
	"strings"
)

// GetAspectRatio returns the arguments for aspect ratio
func (a *Args) GetAspectRatio() []string {
	if a.output.resolution != "" {
		resolution := strings.Split(a.output.resolution, "x")
		if len(resolution) != 0 {
			width, _ := strconv.ParseFloat(resolution[0], 64)
			height, _ := strconv.ParseFloat(resolution[1], 64)
			return []string{"-aspect", fmt.Sprintf("%f", width/height)}
		}
	}
	if a.output.aspectRatio != "" {
		return []string{"-aspect", a.output.aspectRatio}
	}
	return nil
}

// GetConstantRateFactor gets the constant rate factor (CRF) for video encoding
func (a *Args) GetConstantRateFactor() []string {
	if a.output.constantRateFactor != -1 {
		return []string{"-crf", fmt.Sprintf("%d", a.output.constantRateFactor)}
	}

	return nil
}

// GetVideoBitRate returns returns the arguments for video bit rate
func (a *Args) GetVideoBitRate() []string {
	if a.output.videoBitRate != 0 {
		return []string{"-b:v", fmt.Sprintf("%d", a.output.videoBitRate)}
	}
	return nil
}

// GetVideoMaxBitrate returns the arguments for video max bit rate
func (a *Args) GetVideoMaxBitrate() []string {
	if a.output.videoMaxBitrate != 0 {
		return []string{"-maxrate", fmt.Sprintf("%dk", a.output.videoMaxBitrate)}
	}
	return nil
}

// GetVideoMinBitrate returns the arguments for video min bit rate
func (a *Args) GetVideoMinBitrate() []string {
	if a.output.videoMinBitrate != 0 {
		return []string{"-minrate", fmt.Sprintf("%dk", a.output.videoMinBitrate)}
	}
	return nil
}

// GetVideoBitRateTolerance returns the arguments for video bit rate tolerance
func (a *Args) GetVideoBitRateTolerance() []string {
	if a.output.videoBitRateTolerance != 0 {
		return []string{"-bt", fmt.Sprintf("%dk", a.output.videoBitRateTolerance)}
	}
	return nil
}

// GetVideoCodec returns the arguments for video codec
func (a *Args) GetVideoCodec() []string {
	if a.output.videoCodec != "" {
		return []string{"-c:v", a.output.videoCodec}
	}
	return nil
}

// GetVFrames returns the arguments for vframes
func (a *Args) GetVFrames() []string {
	if a.output.vFrames != 0 {
		return []string{"-frames:v", fmt.Sprintf("%d", a.output.vFrames)}
	}
	return nil
}

// GetFrameRate returns the arguments for frame rate
func (a *Args) GetFrameRate() []string {
	if a.output.frameRate != 0 {
		return []string{"-r", fmt.Sprintf("%d", a.output.frameRate)}
	}
	return nil
}

// GetAudioRate returns the arguments for audio rate
func (a *Args) GetAudioRate() []string {
	if a.output.audioRate != 0 {
		return []string{"-ar", fmt.Sprintf("%d", a.output.audioRate)}
	}
	return nil
}

// GetAudioBitrate returns the arguments for bitrate
func (a *Args) GetAudioBitrate() []string {
	if a.output.audioBitrate != 0 {
		return []string{"-b:a", fmt.Sprintf("%d", a.output.audioBitrate)}
	}
	return nil
}

// GetPreset returns the preset
func (a *Args) GetPreset() []string {
	if a.output.preset != "" {
		return []string{"-preset", a.output.preset}
	}
	return nil
}

// GetBufferSize returns the buffer size
func (a *Args) GetBufferSize() []string {
	if a.output.bufferSize != 0 {
		return []string{"-bufsize", fmt.Sprintf("%dk", a.output.bufferSize)}
	}
	return nil
}

// GetPixelFormat returns the pixel format
func (a *Args) GetPixelFormat() []string {
	if a.output.pixelFormat != "" {
		return []string{"-pix_fmt", a.output.pixelFormat}
	}
	return nil
}

// GetKeyframeInterval returns the key frame interval
func (a *Args) GetKeyframeInterval() []string {
	if a.output.keyframeInterval != 0 {
		return []string{"-g", fmt.Sprintf("%d", a.output.keyframeInterval)}
	}
	return nil
}

// GetAudioCodec returns the audio codec
func (a *Args) GetAudioCodec() []string {
	if a.output.audioCodec != "" {
		return []string{"-c:a", a.output.audioCodec}
	}
	return nil
}

// GetAudioChannels returns the audio channels
func (a *Args) GetAudioChannels() []string {
	if a.output.audioChannels != 0 {
		return []string{"-ac", fmt.Sprintf("%d", a.output.audioChannels)}
	}
	return nil
}

// GetFormat returns the output format
func (a *Args) GetFormat() []string {
	if a.output.format != "" {
		return []string{"-f", a.output.format}
	}
	return nil
}

// GetQuality returns the quality
func (a *Args) GetQuality() []string {
	if a.output.quality != 0 {
		return []string{"-crf", fmt.Sprintf("%d", a.output.quality)}
	}
	return nil
}

// GetOutputPath returns the output path
func (a *Args) GetOutputPath() []string {
	if a.output.outputPath != "" {
		return []string{a.output.outputPath}
	}
	return nil
}

// GetPipeOutput returns whether or not ffmpeg is set to receive piped output
func (a *Args) GetPipeOutput() []string {
	if a.output.pipeOutput != false {
		return []string{"pipe:1"}
	}

	return nil
}

// GetOverwrite returns whether or not FFmpeg is set to pipe its output
func (a *Args) GetOverwrite() []string {
	if a.output.overwrite != false {
		return []string{"-y"}
	}

	return nil
}

// GetOutputOptions returns the additional output options
func (a *Args) GetOutputOptions() []string {
	return a.output.outputOptions
}
