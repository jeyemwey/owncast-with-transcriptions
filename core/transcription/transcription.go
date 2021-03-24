package transcription

import (
  "encoding/json"
  "fmt"
  log "github.com/sirupsen/logrus"
  "sync"
  "time"

  "github.com/owncast/owncast/core/data"
)

var (
  subtitlesMut               sync.RWMutex
  subtitleSegments           map[string]*WebVttFile
  numberOfChoppedSegmentsMut sync.RWMutex
  numberOfChoppedSegments    int32
)

var UsedTranscriptionService = GetInstanceOfGoogleTranscriptionService()

func SetupTranscription() error {
  UsedTranscriptionService.SetTranscriptionReceiver(transcriptionReceiver)

  subtitleSegments = make(map[string]*WebVttFile)

  streamConnectedOnce = new(sync.Once)

  return nil
}

// tr: Recognition from the Transcription Service
// cr: One chopped recognition
// crs: Array of chopped recognitions
func transcriptionReceiver(tr Recognition) {
  subtitlesMut.Lock()
  defer subtitlesMut.Unlock()
  numberOfChoppedSegmentsMut.Lock()
  defer numberOfChoppedSegmentsMut.Unlock()

  streamConnectedOnce.Do(func() {
    webVttRenderTicker = WebVttRenderScheduler()
    recognitionLatency = time.Since(streamConnectedTime)
    numberOfChoppedSegments = 1
  })

  tr.Number = numberOfChoppedSegments

  crs := tr.Chop(data.GetStreamLatencyLevel().SecondsPerSegment)

  for _, cr := range crs {
    if _, ok := subtitleSegments[dToStr(cr.SegmentStart)]; !ok {
      subtitleSegments[dToStr(cr.SegmentStart)] = &WebVttFile{SegmentStart: tr.SegmentStart}
    }

    subtitleSegments[dToStr(cr.SegmentStart)].Recognitions =
      append(subtitleSegments[dToStr(cr.SegmentStart)].Recognitions, cr)
  }
  numberOfChoppedSegments += int32(len(crs))
}

func dToStr(t time.Duration) string {
  return fmt.Sprintf("%s", t)
}

func WebVttRenderScheduler() *time.Ticker {
  latencyLevel := data.GetStreamLatencyLevel()
  videoSegmentLength := time.Second * time.Duration(latencyLevel.SecondsPerSegment)
  ticker := time.NewTicker(videoSegmentLength)

  go func() {
    renderCounter := 0
    filesforPlaylist := []*WebVttFile{}

    for {
      select {
      case <-ticker.C:
        go func(renderCounter int) {
          searchString := dToStr(videoSegmentLength * time.Duration(renderCounter))
          log.Info("Looking for: " + searchString + " in:")
          logJson(subtitleSegments)

          webvttFile, ok := subtitleSegments[searchString]

          if !ok {
            log.Info("Failed to find subtitleSegment, creating empty")
            webvttFile = new(WebVttFile)
          }

          webvttFile.FileName = fmt.Sprintf("subtitles%d.webvtt", renderCounter)
          webvttFile.FileContent = webvttFile.AsWebVtt()
          webvttFile.Transmit()

          if len(filesforPlaylist) > 2 {
            filesforPlaylist = filesforPlaylist[1:] // Drop oldest item and
          }
          filesforPlaylist = append(filesforPlaylist, webvttFile) // add the latest item to the playlist

          playlist := makePlaylistFromWebVttFiles(filesforPlaylist, renderCounter)
          logJson(playlist)
          playlist.Transmit()
        }(renderCounter)
        renderCounter += 1
      }
    }
  }()

  return ticker
}

func logJson(u interface{}) {
 b, err := json.Marshal(u)
 if err != nil {
   log.Error(err)
   return
 }
 log.Info(string(b))
}
