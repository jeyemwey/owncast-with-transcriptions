package transcription

//import log "github.com/sirupsen/logrus"
//
//func GetInstanceOfAzureTranscriptionService() TranscriptionService {
//  log.Panic("Azure was not compiled ")
//  return nil
//}

// =========================================================================
// To use Azure comment out everything above and uncomment everything below.
// Do not uncomment the first line (`package transcription`).
// =========================================================================

import (
  "fmt"
  "github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
  "github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
  log "github.com/sirupsen/logrus"
  "os"
)

type AzureTranscriptionService struct {
  TranscriptionReceiver TranscriptionReceiver
  audioConfig           *audio.AudioConfig
  err                   error
  config                *speech.SpeechConfig
  speechRecognizer      *speech.SpeechRecognizer
  task                  chan speech.SpeechRecognitionOutcome
  stream                *audio.PushAudioInputStream
  rawPcmBuffer          *os.File
}

func (a *AzureTranscriptionService) SetConnected() {

  a.stream, a.err = audio.CreatePushAudioInputStream()
  log.Info("Started Audio Stream")
  if a.err != nil {
    log.Error("Got an error: ", a.err)
    return
  }
  if a.stream == nil {
    log.Error("Stream is still nil!")
  }

  a.audioConfig, a.err = audio.NewAudioConfigFromStreamInput(a.stream)
  if a.err != nil {
    log.Error("Got an error: ", a.err)
    return
  }
  a.config, a.err = speech.NewSpeechConfigFromSubscription(Config.Azure.Key, Config.Azure.Region)
  if a.err != nil {
    log.Error("Got an error: ", a.err)
    return
  }
  a.err = a.config.SetSpeechRecognitionLanguage(Config.Language)
  if a.err != nil {
    log.Error("Got an error: ", a.err)
    return
  }
  a.speechRecognizer, a.err = speech.NewSpeechRecognizerFromConfig(a.config, a.audioConfig)
  if a.err != nil {
    log.Error("Got an error: ", a.err)
    return
  }
  a.speechRecognizer.SessionStarted(logSessionWithAnnotation("Started Session with id ="))
  a.speechRecognizer.SessionStopped(logSessionWithAnnotation("Session Stopped with id ="))
  a.speechRecognizer.Canceled(func(event speech.SpeechRecognitionCanceledEventArgs) {
    defer event.Close()
    logJson(event)
  })

  if deliverymethod == "websockets" {
    a.speechRecognizer.Recognizing(func(event speech.SpeechRecognitionEventArgs) {
      defer event.Close()

      SendTranscriptionToWebsocket(event.Result.Text, event.Result.Offset.Nanoseconds())
    })
  }

  if deliverymethod == "webvtt" {
    a.speechRecognizer.Recognized(func(event speech.SpeechRecognitionEventArgs) {
      defer event.Close()

      log.Info("Got a recognition:")
      logJson(event.Result)

      a.TranscriptionReceiver(Recognition{
        Text:     event.Result.Text,
        Begin:    event.Result.Offset,
        End:      event.Result.Offset + event.Result.Duration,
        Duration: event.Result.Duration,
      })
    })
  }

  errChan := a.speechRecognizer.StartContinuousRecognitionAsync()
  go func() {
    for {
      err := <-errChan
      if err == nil {
        continue
      }
      log.Error("Continuous Recognition error", err)
    }
  }()
}

func logSessionWithAnnotation(s string) speech.SessionEventHandler {
  return func(event speech.SessionEventArgs) {
    defer event.Close()
    fmt.Println(s, event.SessionID)
  }
}

func (a *AzureTranscriptionService) SetDisconnected() {
  defer a.audioConfig.Close()
  defer a.config.Close()
  defer a.speechRecognizer.Close()
}

func (a *AzureTranscriptionService) SetTranscriptionReceiver(receiver TranscriptionReceiver) {
  a.TranscriptionReceiver = receiver
}

func (a *AzureTranscriptionService) HandlePcmData(pcmBytes []byte) {
  if a.stream == nil {
    panic("illegal state: a.stream is nil")
  }

  // Write writes the audio data specified by making an internal copy of the
  // data. Note: The dataBuffer should not contain any audio header
  err := a.stream.Write(pcmBytes)
  if err != nil {
    log.Error("unable to write wav to stream with error:", err)
  }
}

var ats *AzureTranscriptionService

// Because this is the only way that I can tell Go to explicitly typecheck me
// with an inferring interfaces.
func GetInstanceOfAzureTranscriptionService() *AzureTranscriptionService {
  if ats == nil {
    ats = &AzureTranscriptionService{}
  }

  return ats
}
