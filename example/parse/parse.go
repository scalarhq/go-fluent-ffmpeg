package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	cmd := fluentffmpeg.NewCommand("")
	cmd.InputPath("/path/to/output/file")
	cmd.OutputPath("/path/to/input/file")
	cmd.ConstantRateFactor(22)
	cmd.Overwrite(true)
	cmd.Options("-progress", "pipe:2", "-v", "quiet")

	c := cmd.Build()

	outReader, err := c.StderrPipe()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer outReader.Close()

	scanner := bufio.NewScanner(outReader)
	scanner.Split(fluentffmpeg.SplitProgress)

	go func() {
		for scanner.Scan() {
			p, err := fluentffmpeg.ParseProgress(bytes.NewReader(scanner.Bytes()))
			if err != nil {
				log.Fatal(err)
			}
			if p.Status != "" {
				fmt.Printf("Status: %s Frame: %d Time: %d\n", p.Status, p.Frame, p.OutTime.Milliseconds())
			}
		}
	}()

	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
