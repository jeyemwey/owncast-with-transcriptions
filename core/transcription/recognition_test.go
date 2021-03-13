package transcription

import (
  "reflect"
  "testing"
  "time"
)

func TestRecognition_Chop_just_in_first_segment(t *testing.T) {
  r := Recognition{"", 1, dur("1s"), dur("3s"), dur("2s"), 0}

  cr := r.Chop(4)

  AssertEqual(t, cr, []ChoppedRecognition{
    {"", 1, dur("1s"), dur("3s"), dur("2s"), dur("0s")}})
}

func TestRecognition_Chop_just_in_nth_segment(t *testing.T) {
  r := Recognition{"", 1, dur("9s"), dur("11s"), dur("2s"), 0}

  cr := r.Chop(4)

  AssertEqual(t, cr, []ChoppedRecognition{
    {"", 1, dur("9s"), dur("11s"), dur("2s"), dur("8s")}})
}

func TestRecognition_Chop_into_first_and_second_segment(t *testing.T) {
  r := Recognition{"", 1, dur("3s"), dur("5s"), dur("2s"), 0}

  cr := r.Chop(4)

  AssertEqual(t, cr, []ChoppedRecognition{
    {"", 1, dur("3s"), dur("4s"), dur("1s"), dur("0s")},
    {"", 2, dur("4s"), dur("5s"), dur("1s"), dur("4s")}})
}

func TestRecognition_Chop_into_first_second_and_third_segment(t *testing.T) {
  r := Recognition{"", 1, dur("3s"), dur("9s"), dur("2s"), 0}

  cr := r.Chop(4)

  AssertEqual(t, cr, []ChoppedRecognition{
    {"", 1, dur("3s"), dur("4s"), dur("1s"), dur("0s")},
    {"", 2, dur("4s"), dur("8s"), dur("4s"), dur("4s")},
    {"", 3, dur("8s"), dur("9s"), dur("1s"), dur("8s")}})
}

func TestRecognition_Chop_into_second_third_and_fourth_segment(t *testing.T) {
  r := Recognition{"", 1, dur("6s"), dur("14s"), dur("2s"), 0}

  cr := r.Chop(4)

  AssertEqual(t, cr, []ChoppedRecognition{
    {"", 1, dur("6s"), dur("8s"), dur("2s"), dur("4s")},
    {"", 2, dur("8s"), dur("12s"), dur("4s"), dur("8s")},
    {"", 3, dur("12s"), dur("14s"), dur("2s"), dur("12s")}})
}

func dur(s string) time.Duration {
  t, _ := time.ParseDuration(s)
  return t
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, received interface{}, expected interface{}) {
  if reflect.DeepEqual(received, expected) {
    return
  }
  // debug.PrintStack()
  t.Errorf("Received %v (type %v), expected %v (type %v)", received, reflect.TypeOf(received), expected, reflect.TypeOf(expected))
}
