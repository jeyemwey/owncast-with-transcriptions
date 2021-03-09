package transcription

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func RewriteMasterFile(path string) {
  file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Error("Cannot open HLS master file", path, err)
    return
  }
  defer file.Close()

  webVTTStaticText := "\n#EXT-X-MEDIA:" +
    "TYPE=SUBTITLES," +
    "GROUP-ID=\"subs\"," +
    "NAME=\"English\"," +
    "DEFAULT=NO," +
    "FORCED=NO," +
    "URI=\"subtitles.m3u8\"," +
    "LANGUAGE=\"en\"\n"

  if _, err := file.WriteString(webVTTStaticText); err != nil {
    log.Error("Cannot append HLS master file", path, err)
  }
}