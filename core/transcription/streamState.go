package transcription

func SetConnected() {
  go StartAudioTranscodingForTranscriptionService()
  go UsedTranscriptionService.SetConnected()
}

func SetDisconnected(){
  go UsedTranscriptionService.SetDisconnected()
}
