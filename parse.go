package fluentffmpeg

import (
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type ProgressStatus string

const (
	// ProgessStatusContinue indicates that FFmpeg will continue processing
	ProgessStatusContinue ProgressStatus = "continue"
	// ProgessStatusComplete indicates that FFmpeg has completed
	ProgessStatusComplete ProgressStatus = "complete"
)

// Progress contains FFmpeg progress information
type Progress struct {
	Status ProgressStatus
	Frame  int64
	FPS    float32
	// The duration of video that has been processed
	OutTime time.Duration
	// Other contains other progress information
	Other map[string]string
}

// ParseProgress reads from `r` and parses progress information. `r` must contain
// valid FFmpeg progress information.
func ParseProgress(r io.Reader) (Progress, error) {
	p := Progress{}
	b, err := io.ReadAll(r)
	if err != nil {
		return p, err
	}
	str := string(b)
	lines := strings.Split(str, "\n")
	if len(lines) == 0 {
		return p, nil
	}
	for _, line := range lines {
		if strings.Contains(line, "=") {
			keyVal := strings.Split(line, "=")
			if len(keyVal) >= 2 {
				// Trim key/val because they may contain leading spaces
				err = parseProgressKeyValue(&p, strings.Trim(keyVal[0], " "), strings.Trim(keyVal[1], " "))
				if err != nil {
					return p, errors.Wrap(err, "failed to parse progress")
				}
			}
		}
	}
	return p, nil
}

func parseProgressKeyValue(p *Progress, key string, value string) error {
	switch key {
	case "progress":
		p.Status = ProgressStatus(value)
	case "frame":
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		p.Frame = i
	case "fps":
		i, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		p.FPS = float32(i)
	case "out_time_ms":
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		p.OutTime = time.Microsecond * time.Duration(i)
	default:
		if p.Other == nil {
			p.Other = make(map[string]string)
		}
		p.Other[key] = value
	}
	return nil
}

// SplitProgress is a bufio.SplitFunc for scanning FFmpeg progress data
func SplitProgress(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	str := string(data)
	// "progress=continue" or "progress=end" denotes the end
	// of one FFmpeg progress update
	i := strings.Index(str, "progress=continue")
	offset := len("progress=continue")
	if i == -1 {
		i = strings.Index(str, "progress=end")
		offset = len("progress=end")
	}
	if i >= 0 {
		return i + offset, data[0 : i+offset], nil
	}
	// If we're at EOF, we have a final, non-terminated FFmpeg progress. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
