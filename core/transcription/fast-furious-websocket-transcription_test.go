package transcription

import "testing"

func Test_lastSentence(t *testing.T) {
  tests := []struct {
    name string
    in string
    want string
  }{
    {name: "No punct.", in: "This is a sentence", want: "This is a sentence"},
    {name: "One punct.", in: "This is a sentence with a question mark?", want: "This is a sentence with a question mark?"},
    {name: "Two sentences", in: "This is a sentence. Another sentence.", want: "Another sentence."},
    {name: "Last sentence has no punct.", in: "This is a sentence. This is another sentence without dot", want: "This is another sentence without dot"},
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := lastSentence(tt.in); got != tt.want {
        t.Errorf("lastSentence() = %v, want %v", got, tt.want)
      }
    })
  }
}
