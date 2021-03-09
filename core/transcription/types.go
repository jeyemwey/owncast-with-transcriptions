package transcription

import "time"

type Recognition struct {
  Text string

  Begin    time.Duration // Relative to stream start
  End      time.Duration // Relative to stream start
  Duration time.Duration // Relative to begin
}

type TranscriptionService interface {
  HandlePcmData(pcmData []byte)
  SetTranscriptionReceiver(receiver TranscriptionReceiver)
  SetConnected()
  SetDisconnected()
}

type TranscriptionReceiver func(recognition Recognition)


type WebVttFile struct {
 Name        string
 FileContent string
}
