package transcription

func SetConnected() {
  go StartAudioTranscodingForTranscriptionService()
}

func SetDisconnected(){
}
