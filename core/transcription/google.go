package transcription

import (
	"context"
	"io"

	speech "cloud.google.com/go/speech/apiv1"
	log "github.com/sirupsen/logrus"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type GoogleTranscriptionService struct {
  TranscriptionReceiver TranscriptionReceiver
  ctx                   context.Context
  err                   error
  client                *speech.Client
  stream                speechpb.Speech_StreamingRecognizeClient
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
  g.ctx = context.Background()

  g.client, g.err = speech.NewClient(g.ctx)
  if g.err != nil {
    log.Fatal(g.err)
  }
  g.stream, g.err = g.client.StreamingRecognize(g.ctx)
  if g.err != nil {
    log.Fatal(g.err)
  }
  log.Info("Started StreamingRecognize with Google")
  // Send the initial configuration message.
  if g.err = g.stream.Send(&speechpb.StreamingRecognizeRequest{
    StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{

      StreamingConfig: &speechpb.StreamingRecognitionConfig{
          Config: &speechpb.RecognitionConfig{
          Encoding:                   speechpb.RecognitionConfig_LINEAR16,
          SampleRateHertz:            16000,
          LanguageCode:               gcpLanguage,
          EnableAutomaticPunctuation: true,
          MaxAlternatives: 8,
        },
        InterimResults: true,
      },
    },
  }); g.err != nil {
    log.Fatal(g.err)
  }

  // If anything is broken, we don't need to handle the receivers.
  if g.err != nil {
    return
  }

  //lastResultEndTime := 0 * time.Nanosecond
  log.Info("Started looking for responses from Google")

  for {
    resp, err := g.stream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatalf("Cannot stream results: %v", err)
    }

    if err := resp.Error; err != nil {
      // Workaround while the API doesn't give a more informative error.
      if err.Code == 3 || err.Code == 11 {
        log.Warn("Speech recognition request exceeded limit of 300 seconds.")
      }
      log.Fatalf("Could not recognize: %v", err)
    }
    for _, result := range resp.Results {

      if result.Stability < 0.6 || result.IsFinal {
        continue
      }

      //log.Infof("Got a result: %+v\n", result)

      bestAlternative := result.GetAlternatives()[0]
      for _, a := range result.GetAlternatives() {
        if a.Confidence > bestAlternative.Confidence {
          bestAlternative = a
        }
      }

      SendTranscriptionToWebsocket(bestAlternative.Transcript, result.ResultEndTime.AsDuration().Nanoseconds())

      //endTime := result.GetResultEndTime().AsDuration()
      //recognition := Recognition{
      //  Text:         result.GetAlternatives()[0].Transcript,
      //  Begin:        lastResultEndTime,
      //  End:          endTime,
      //  Duration:     endTime - lastResultEndTime,
      //}
      //lastResultEndTime = endTime

      //g.TranscriptionReceiver(recognition)
    }
  }

}

func (g *GoogleTranscriptionService) SetDisconnected() {
  if err := g.stream.CloseSend(); err != nil {
    log.Fatalf("Could not close stream: %v", err)
  }
}

func (g *GoogleTranscriptionService) HandlePcmData(pcmData []byte) {
  if g.stream == nil {
    log.Error("g.stream is not ready yet, possible NPE.")
    return
  }

  if err := g.stream.Send(&speechpb.StreamingRecognizeRequest{
    StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
      AudioContent: pcmData,
    },
  }); err != nil {
    log.Printf("Could not send audio: %v", err)
  }
}
