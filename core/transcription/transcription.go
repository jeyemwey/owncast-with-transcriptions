package transcription

import (
  "sync"
)

var (
  subtitlesMut sync.RWMutex
  subtitles    []Recognition
)

var UsedTranscriptionService = GetInstanceOfAzureTranscriptionService()


func SetupTranscription() error {
  UsedTranscriptionService.SetTranscriptionReceiver(transcriptionReceiver)

  return nil
}

func transcriptionReceiver(r Recognition) {
  subtitlesMut.Lock()
  defer subtitlesMut.Unlock()

  subtitles = append(subtitles, r)
}
