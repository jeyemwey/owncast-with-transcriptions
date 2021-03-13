package transcription

import (
  "time"
)

func (r Recognition) Chop(secondsPerSegment int) []ChoppedRecognition {
  segmentLength := (time.Duration(secondsPerSegment)) * time.Second
  iBegin := r.Begin.Truncate(segmentLength)
  iEnd := r.End.Truncate(segmentLength)

  number := r.Number

  var chops []ChoppedRecognition
  for i := iBegin; i <= iEnd; i = i + segmentLength {
    c := ChoppedRecognition{
      Text:         r.Text,
      Number:       number,
      Begin:        maxDuration(i, r.Begin),
      End:          minDuration(i+segmentLength, r.End),
      SegmentStart: i,
    }
    c.Duration = c.End - c.Begin
    chops = append(chops, c)

    number += 1
  }

  return chops
}

func minDuration(a, b time.Duration) time.Duration {
  if a <= b {
    return a
  }
  return b
}

func maxDuration(a, b time.Duration) time.Duration {
  if a >= b {
    return a
  }
  return b
}
