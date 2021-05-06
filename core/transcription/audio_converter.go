package transcription

import (
	"fmt"
  "github.com/owncast/owncast/core/timing"
  "io"
	"os/exec"
	"strings"
  "sync"
  "time"

  "github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/utils"
	log "github.com/sirupsen/logrus"
)

func StartAudioTranscodingForTranscriptionService() {
  ffmpegFlags := []string{
    utils.ValidatedFfmpegPath(data.GetFfMpegPath()),
    "-i", utils.GetTemporaryTranscriptionPipePath(),
    "-vn",
    "-c:a pcm_s16le",
    "-f s16le",
    "-ac 1",
    "-ar 16000",
    "-",
  }

  ffmpegCmd := strings.Join(ffmpegFlags, " ")
  log.Debug(ffmpegCmd)

  cmd := exec.Command("sh", "-c", ffmpegCmd)
  // stolen from https://stackoverflow.com/a/43602656
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Error(err)
    return
  }

  timing.D.TranscodingStarted = time.Now()

  if err = cmd.Start(); err != nil {
    log.Error(err)
    return
  }

  var once sync.Once

  pcmChunk := make([]byte, pcmChunkLen)
  for {
    pcmLength, err := stdout.Read(pcmChunk)

    go once.Do(func() {
        timing.D.TranscodingReturnedFirstResult = time.Now()
      })

    if pcmLength > 0 {
      validPcmBytes := pcmChunk[:pcmLength]
      Config.UsedTranscriptionService.HandlePcmData(validPcmBytes)
    }

    if err != nil {
      if err == io.EOF {
        break
      }
      fmt.Printf("Error = %v\n", err)
      continue
    }
  }

  if err := cmd.Wait(); err != nil {
    fmt.Printf("Wait command error: %v\n", err)
  }
}
