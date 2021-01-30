package main

import (
	"fmt"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	data, _ := fluentffmpeg.Probe("./video.avi")

	fmt.Println(data["format"])
}
