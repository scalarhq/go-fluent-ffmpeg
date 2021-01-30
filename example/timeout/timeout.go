package main

import (
	"fmt"
	"time"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	done := make(chan error, 1)
	cmd := fluentffmpeg.NewCommand("").
		InputPath("./video.avi").
		OutputFormat("mp4").
		OutputPath("./video.mp4").
		Overwrite(true).
		Build()
	cmd.Start()

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("Timed out")
		cmd.Process.Kill()
	case <-done:
	}
}
