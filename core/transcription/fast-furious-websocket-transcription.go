package transcription

import (
  "fmt"
  "github.com/owncast/owncast/core/chat"
  "github.com/owncast/owncast/models"
  "strings"
  "sync"
  "time"
)

var lastWebsocketPush time.Time
var lwp_mx sync.Mutex

// This function will send a Websocket message with type "SUBTITLE" to every
// client that is connected. To reduce the network load, we need a certain
// duration between the pushs. Every request to push that is too close to the
// last request is dropped to keep the latency down.
func SendTranscriptionToWebsocket(body string, ns_since_begin int64) {
  lwp_mx.Lock()
  defer lwp_mx.Unlock()
  if time.Since(lastWebsocketPush) < 200*time.Millisecond {
    return
  }
  lastWebsocketPush = time.Now()

  // log.Infof("timeSinceBegin: %v; body: %s", ns_since_begin, body)

  go chat.SendMessage(models.ChatEvent{
    Body:           strings.TrimSpace(lastSentence(body)),
    ID:             fmt.Sprintf("subtitle-%d", time.Now().Unix()),
    MessageType:    "SUBTITLE",
    Timestamp:      time.Now(),
    TimeSinceBegin: ns_since_begin,
    Ephemeral:      true,
  })
}

// This function will only return the last sentence of a given chain of sentences
// and reduce the amount of content presented to the viewer.
//
// TODO: We need to reduce the amount of content further, if the sentences are
// too long or punctuation is missing in the string. While this is a story for
// another day, this would be the place to put it.
func lastSentence(body string) string {
  lastPunctation := strings.LastIndexAny(body, ",.?!")
  if lastPunctation != len(body) { // Leave the last char if it is a punctuation
    // Search everywhere except for the last char
    lastPunctation = strings.LastIndexAny(body[:len(body)-1], ",.?!")
  }

  // No punctuation found
  if lastPunctation == -1 { // No punctuation found
    return body
  }

  return body[lastPunctation+2:] // remove the dot and the space after it, too
}
