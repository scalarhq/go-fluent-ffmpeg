package fluentffmpeg

// Args contains the input and output args set for FFmpeg
type Args struct {
	input         inputArgs
	output        outputArgs
	globalOptions []string
}

type inputArgs struct {
	inputPath            string
	pipeInput            bool
	fromFormat           string
	nativeFramerateInput bool `getter:"none"`
	inputOptions         []string
}

type outputArgs struct {
	outputPath            string
	format                string
	pipeOutput            bool
	overwrite             bool
	resolution            string `getter:"none"`
	aspectRatio           string
	pixelFormat           string
	quality               int
	preset                string
	bufferSize            int
	audioBitrate          int
	audioChannels         int
	keyframeInterval      int
	audioCodec            string
	constantRateFactor    int
	videoBitRate          int
	videoBitRateTolerance int
	videoMaxBitrate       int
	videoMinBitrate       int
	videoCodec            string
	vFrames               int
	frameRate             int
	audioRate             int
	outputOptions         []string
}

func (a *Args) GetOptions() []string {
	return a.globalOptions
}
