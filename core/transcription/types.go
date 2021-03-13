package transcription

import "time"

type Recognition struct {
  Text   string
  Number int32

  Begin        time.Duration // Relative to stream start
  End          time.Duration // Relative to stream start
  Duration     time.Duration // Relative to begin
  SegmentStart time.Duration // Relative to stream start, begin of the index
}

// A ChoppedRecognition is a Recognition object with cr.End - cr.Begin <=
// config.videoSegmentDuration. Only ChoppedRecognitions can get fed into a WebVttFile.
type ChoppedRecognition Recognition

type File struct {
  FileName    string
  FileContent string
}

type WebVttFile struct {
  File

  Recognitions []ChoppedRecognition
  SegmentStart time.Duration
}

type TranscriptionService interface {
  HandlePcmData(pcmData []byte)
  SetTranscriptionReceiver(receiver TranscriptionReceiver)
  SetConnected()
  SetDisconnected()
}

type TranscriptionReceiver func(recognition Recognition)
