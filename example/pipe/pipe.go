package main

import (
	"net/http"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
)

func main() {
	http.HandleFunc("/ffmpeg", handle)

	http.ListenAndServe(":5000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, _ := r.FormFile("video")

	fluentffmpeg.
		NewCommand("").
		PipeInput(file).
		OutputFormat("flv").
		PipeOutput(w).
		Run()
}
