package transcription

import (
  "fmt"
  "github.com/owncast/owncast/core/chat"
  "github.com/owncast/owncast/models"
  "strings"
  "time"
)

func SendTranscriptionToWebsocket(body string) {
  body = lastSentence(body)

  chat.SendMessage(models.ChatEvent{
    Body:        body,
    ID:          fmt.Sprintf("subtitle-%d", time.Now().Unix()),
    MessageType: webSocketSubtitlesType,
    Timestamp:   time.Now(),
    Ephemeral:   true,
  })
}

func lastSentence(body string) string {
  lastPunctation := strings.LastIndexAny(body, ",.?!")
  if lastPunctation != len(body) { // Last punctuation
    lastPunctation = strings.LastIndexAny(body[:len(body)-1], ",.?!")
  }

  if lastPunctation == -1 { // No punctuation found
    return body
  }

  return body[lastPunctation + 2:] // remove the dot and the space after it
}
