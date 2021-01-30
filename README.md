# Go Fluent FFmpeg

A Go version of [node-fluent-ffmpeg](https://github.com/fluent-ffmpeg/node-fluent-ffmpeg).

## Installation
`go get -u github.com/modfy/go-fluent-ffmpeg`

### Requirements
You will need FFmpeg installed on your machine, or you can specify a path to a binary:

```go
// Provide an empty string to use default FFmpeg path
cmd := fluentffmpeg.NewCommand("")

// Specify a path
cmd = fluentffmpeg.NewCommand("/path/to/ffmpeg/binary")
```

## Quick Start

Create and run commands using an API similar to node-fluent-ffmpeg:

```go
err := fluentffmpeg.NewCommand(""). 
		InputPath("/path/to/video.avi").
		OutputFormat("mp4").
		OutputPath("/path/to/video.mp4").
		Run()
```

If you want to view the errors/logs returned from FFmpeg, provide an io.Writer to receive the data. 
```go
buf := &bytes.Buffer{}
err := fluentffmpeg.NewCommand("").
		InputPath("./video.avi").
		OutputFormat("mp4").
		OutputPath("./video.mp4").
		Overwrite(true).
		OutputLogs(buf). // provide a io.Writer
        Run()

out, _ := ioutil.ReadAll(buf) // read logs
fmt.Println(string(out))
```

You can also get the command in the form of an [exec.Cmd](https://golang.org/pkg/os/exec/#Cmd) struct, with which you can have better control over the running process. For example, you can conditionally kill the FFmpeg command:

```go
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
```

## Credits

This repo was inspired by [node-fluent-ffmpeg](https://github.com/fluent-ffmpeg/node-fluent-ffmpeg) and was built upon the work done by [@bitcodr](https://github.com/bitcodr/) in the https://github.com/bitcodr/gompeg

## Managed Version

You can deploy this codebase yourself or you an entirely managed api from the creators at https://api.modfy.video