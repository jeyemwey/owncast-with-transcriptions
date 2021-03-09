package transcription

type GoogleTranscriptionService struct {
  TranscriptionReceiver TranscriptionReceiver
}

var gts *GoogleTranscriptionService

// Because this is the only way that I can tell Go to explicitly typecheck me
// with an inferring interfaces.
func GetInstanceOfGoogleTranscriptionService() *GoogleTranscriptionService {
  if gts == nil {
    gts = &GoogleTranscriptionService{}
  }

  return gts
}

func (g *GoogleTranscriptionService) SetTranscriptionReceiver(receiver TranscriptionReceiver) {
  g.TranscriptionReceiver = receiver
}

func (g *GoogleTranscriptionService) SetConnected() {
  panic("implement me")
}

func (g *GoogleTranscriptionService) SetDisconnected() {
  panic("implement me")
}

func (g *GoogleTranscriptionService) HandlePcmData(pcmData []byte) {
  panic("implement me")
}
