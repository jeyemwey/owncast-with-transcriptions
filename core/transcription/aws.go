package transcription

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  tss "github.com/aws/aws-sdk-go/service/transcribestreamingservice"
  log "github.com/sirupsen/logrus"
  "io"
  "time"
)

type AwsTranscriptionService struct {
  TranscriptionReceiver TranscriptionReceiver
  sess                  *session.Session
  err                   error
  client                *tss.TranscribeStreamingService
  resp                  *tss.StartStreamTranscriptionOutput
  stream                *tss.StartStreamTranscriptionEventStream
  pcmReader             *io.PipeReader
  pcmWriter             *io.PipeWriter
  streamContext         aws.Context
}

var awsts *AwsTranscriptionService

func GetInstanceOfAwsTranscriptionService() *AwsTranscriptionService {
  if awsts == nil {
    awsts = &AwsTranscriptionService{}
  }

  return awsts
}

func (a *AwsTranscriptionService) SetConnected()  {
  a.sess, a.err = session.NewSession(&aws.Config{
    Region:      aws.String("us-west-2"),
    Credentials: credentials.NewSharedCredentials("", "bsc-thesis"),
  })
  if a.err != nil {
    log.Fatalf("failed to load SDK configuration, %v", a.err)
  }

  a.client = tss.New(a.sess)

  a.resp, a.err = a.client.StartStreamTranscription(&tss.StartStreamTranscriptionInput{
    LanguageCode:         aws.String(awsLanguage),
    MediaEncoding:        aws.String(tss.MediaEncodingPcm),
    MediaSampleRateHertz: aws.Int64(16000),
  })
  if a.err != nil {
    log.Fatalf("failed to start streaming, %v", a.err)
  }

  a.stream = a.resp.GetStream()
  defer a.stream.Close()

  a.pcmReader, a.pcmWriter = io.Pipe()
  a.streamContext = aws.BackgroundContext()

  go func() {
    err := tss.StreamAudioFromReader(a.streamContext, a.stream, pcmChunkLen, a.pcmReader)
    if err != nil {
      log.Errorf("Error with StreamAudioFromReader: %v", err)
    }
  }()


  for event := range a.stream.Events() {
    switch e := event.(type) {
    case *tss.TranscriptEvent:
      for _, res := range e.Transcript.Results {
        for _, alt := range res.Alternatives {
          body := aws.StringValue(alt.Transcript)
          ns := FloatSecsToNanoSeconds(res.StartTime)
          //log.Infof("%f / %d: %s", *res.StartTime, ns, body)
          if deliverymethod == "websockets" {
            go SendTranscriptionToWebsocket(body, ns)
          }

          if !*res.IsPartial && deliverymethod == "webvtt" {
            // Only for final results
            transcriptionReceiver(Recognition{
              Text:  body,
              Begin: time.Duration(FloatSecsToNanoSeconds(res.StartTime)) * time.Nanosecond,
              End:   time.Duration(FloatSecsToNanoSeconds(res.EndTime)) * time.Nanosecond,
            })
          }
        }
      }
    default:
      log.Fatalf("unexpected event, %T", event)
    }
  }

  if err := a.stream.Err(); err != nil {
    log.Fatalf("expect no error from stream, got %v", err)
  }
}

func (a *AwsTranscriptionService) SetDisconnected()  {
  a.pcmWriter.Close()
  a.stream.Close()
}

func (a *AwsTranscriptionService) HandlePcmData(b []byte) {
  _, err := a.pcmWriter.Write(b)
  if err != nil {
    log.Errorf("unable to write to writer: %v", err)
  }
}

func (a *AwsTranscriptionService) SetTranscriptionReceiver(receiver TranscriptionReceiver) {
  a.TranscriptionReceiver = receiver
}

func FloatSecsToNanoSeconds(seconds *float64) int64 {
  ns := *seconds * 10e9
  return int64(ns)
}
