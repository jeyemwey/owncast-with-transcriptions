package transcription

import (
  "fmt"
  "strconv"
  "time"
)

func (w WebVttFile) AsWebVtt() (ret string) {
  ret = fmt.Sprintf("WEBVTT - %s\n\n", formatDurationWebVTT(w.SegmentStart))

  for i, r := range w.Recognitions {
    ret += fmt.Sprint(r.Number) + "\n"
    ret = ret + fmt.Sprintf("%s --> %s\n", formatDurationWebVTT(r.Begin), formatDurationWebVTT(r.End))
    ret = ret + r.Text + "\n"

    if i != len(w.Recognitions) {
      ret = ret + "\n"
    }
  }

  return
}

// formatDurationWebVTT formats a .vtt duration
//
// Attribution: Copied from go-astisub which is licensed under the MIT license
// https://github.com/asticode/go-astisub/blob/master/webvtt.go#L275
func formatDurationWebVTT(i time.Duration) string {
  return formatDuration(i, ".", 3)
}

// Attribution: Copied from go-astisub which is licensed under the MIT license
// https://github.com/asticode/go-astisub/blob/master/webvtt.go#L275
func formatDuration(i time.Duration, millisecondSep string, numberOfMillisecondDigits int) (s string) {
  // Parse hours
  var hours = int(i / time.Hour)
  var n = i % time.Hour
  if hours < 10 {
    s += "0"
  }
  s += strconv.Itoa(hours) + ":"

  // Parse minutes
  var minutes = int(n / time.Minute)
  n = i % time.Minute
  if minutes < 10 {
    s += "0"
  }
  s += strconv.Itoa(minutes) + ":"

  // Parse seconds
  var seconds = int(n / time.Second)
  n = i % time.Second
  if seconds < 10 {
    s += "0"
  }
  s += strconv.Itoa(seconds) + millisecondSep

  // Parse milliseconds
  var milliseconds = float64(n/time.Millisecond) / float64(1000)
  s += fmt.Sprintf("%."+strconv.Itoa(numberOfMillisecondDigits)+"f", milliseconds)[2:]
  return
}
