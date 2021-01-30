package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	buf := &bytes.Buffer{}
	err := fluentffmpeg.NewCommand("").
		InputPath("./video.avi").
		OutputFormat("mp4").
		OutputPath("./video.mp4").
		Overwrite(true).
		OutputLogs(buf).
		Run()

	if err != nil {
		log.Fatal(err)
	}

	out, _ := ioutil.ReadAll(buf)
	fmt.Println(string(out))
}
