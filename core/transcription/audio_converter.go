package transcription

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

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

  if err = cmd.Start(); err != nil {
    log.Error(err)
    return
  }

  pcmChunk := make([]byte, pcmChunkLen)
  for {
    pcmLength, err := stdout.Read(pcmChunk)

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
