package timing

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
  OnceMaster sync.Once
  OnceSegment sync.Once
  OnceTrack sync.Once

  D struct {
    // Base time
    Connected time.Time

    // HLS Playlist
    FirstMasterPlaylist time.Time
    FirstTrackPlaylist  time.Time
    FirstTsFile         time.Time

    // Transcription
    TranscodingStarted time.Time
    TranscodingReturnedFirstResult time.Time

    // Cloud
    CloudStarted   time.Time
    CloudConnected time.Time
    CloudResponses []struct {
      Time time.Time
      Text string
    }
  }
)


func CallConnected() {
  time.AfterFunc(1 * time.Minute, func() {
    bytes, err := json.Marshal(D)
    if err != nil {
      log.Error(err)
      return
    }

    if err := ioutil.WriteFile("gcp-de.json", bytes, 0644); err != nil {
      log.Error(err)
    }
  })
}
