package transcription

import "fmt"

func ExampleWebVttFile_AsWebVtt_empty() {
  w := WebVttFile{}
  w.FileName = "subtitles001.webvtt"

  fmt.Print(w.AsWebVtt())

  // Output:
  // WEBVTT - 00:00:00.000
  //
}

func ExampleWebVttFile_AsWebVtt_simple() {
  w := WebVttFile{
    Recognitions: []ChoppedRecognition{{
      Text:         "Hello World",
      Number:       1,
      Begin:        dur("0.5s"),
      End:          dur("2.5s"),
      Duration:     dur("2s"),
      SegmentStart: dur("0s")}},
    SegmentStart: dur("0s"),
  }
  w.FileName = "subtitles001.webvtt"

  fmt.Print(w.AsWebVtt())

  // Output:
  // WEBVTT - 00:00:00.000
  //
  // 1
  // 00:00:00.500 --> 00:00:02.500
  // Hello World
}

func ExampleWebVttFile_AsWebVtt_multiple() {
  w := WebVttFile{
    Recognitions: []ChoppedRecognition{
      {
        Text:         "Hello World",
        Number:       1,
        Begin:        dur("0.5s"),
        End:          dur("2.5s"),
        Duration:     dur("2s"),
        SegmentStart: dur("0s"),
      },
      {
        Text:         "Foo Bar Baz",
        Number:       2,
        Begin:        dur("2.5s"),
        End:          dur("4s"),
        Duration:     dur("1.5s"),
        SegmentStart: dur("0s"),
      },
    },
    SegmentStart: dur("0s"),
  }
  w.FileName = "subtitles001.webvtt"

  fmt.Print(w.AsWebVtt())

  // Output:
  // WEBVTT - 00:00:00.000
  //
  // 1
  // 00:00:00.500 --> 00:00:02.500
  // Hello World
  //
  // 2
  // 00:00:02.500 --> 00:00:04.000
  // Foo Bar Baz
}
