package transcription

import (
  "sync"
  "time"
)

var (
  // The ticker is started once the first recognition is received.
  webVttRenderTicker  *time.Ticker
  streamConnectedTime time.Time
  recognitionLatency  time.Duration
  streamConnectedOnce *sync.Once
)

func SetConnected() {
  go StartAudioTranscodingForTranscriptionService()
  go UsedTranscriptionService.SetConnected()
}

func SetDisconnected() {
  go UsedTranscriptionService.SetDisconnected()
  webVttRenderTicker.Stop()
}
