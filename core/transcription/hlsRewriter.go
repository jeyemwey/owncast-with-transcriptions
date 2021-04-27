package transcription

import (
  "io/ioutil"
  "net/http"
  "strconv"
  "strings"

  "github.com/owncast/owncast/config"
  "github.com/owncast/owncast/core/data"
  log "github.com/sirupsen/logrus"
)

const SUBTITLES_PLAYLIST_FILE = "subtitles.m3u8"

func RewriteMasterFile(path string) {

  if deliverymethod != "webvtt" || !Config.EnableTranscription {
    return
  }

  dat, err := ioutil.ReadFile(path)
  if err != nil {
    log.Errorf("Unable to read master file: %v", err)
    return
  }

  master := string(dat)

  webVTTStaticText := "\n#EXT-X-MEDIA:" +
    "TYPE=SUBTITLES," +
    "GROUP-ID=\"subs\"," +
    "NAME=\"English\"," +
    "DEFAULT=NO," +
    "FORCED=NO," +
    "URI=\"" + SUBTITLES_PLAYLIST_FILE + "\"," +
    "LANGUAGE=\"en\"\n"

  newMaster := strings.ReplaceAll(master, ",CODECS", ",SUBTITLES=\"subs\",CODECS") + webVTTStaticText

  err = ioutil.WriteFile(path, []byte(newMaster), 0644)
  if err != nil {
    log.Errorf("Unable to write master file: %v", err)
  }
}

func makePlaylistFromWebVttFiles(files []*WebVttFile, renderCounter int) (f File) {

  targetDuration := strconv.Itoa(data.GetStreamLatencyLevel().SecondsPerSegment)
  mediaSequence := strconv.Itoa(renderCounter)

  f.FileName = SUBTITLES_PLAYLIST_FILE
  f.FileContent = "#EXTM3U\n" +
    "#EXT-X-TARGETDURATION:" + targetDuration + "\n" +
    "#EXT-X-VERSION:3\n" +
    "#EXT-X-MEDIA-SEQUENCE:" + mediaSequence + "\n"

  for _, file := range files {
    f.FileContent += "\n" +
      "#EXTINF:" + targetDuration + "\n" +
      file.FileName
  }



  return
}

// Transmit the WebVtt file to the FileWriterReceiverService. From there, it
// will be sent to the configured storage provider (currently local or s3).
func (w *File) Transmit() {
  bodyReader := strings.NewReader(w.FileContent)

  req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:"+config.InternalHLSListenerPort+"/"+w.FileName, bodyReader)
  if err != nil {
    log.Error("Cannot send request to Internal HLS Receiver")
  }

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    log.Error("Unable to perform request: ", err)
    return
  }

  if resp.StatusCode != http.StatusOK {
    log.Error("HLS Receiver could not process request.", resp)
  }
}
