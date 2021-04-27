package transcription

import (
  log "github.com/sirupsen/logrus"
  "gopkg.in/yaml.v3"
  "io/ioutil"
  "strings"
)

const pcmChunkLen = 1024

type c struct {
  EnableTranscription bool   `yaml:"enableTranscription"`
  Service  string `yaml:"service"`
  Language string `yaml:"language"`

  Azure struct {
    Region string `yaml:"region"`
    Key    string `yaml:"key"`
  } `yaml:"azure"`

  UsedTranscriptionService TranscriptionService `json:"-" yaml:"-"`
}

var Config c

func init() {
  yamlFile, err := ioutil.ReadFile("transcription.yaml")
  if err != nil {
    log.Info(err)
    log.Info("Transcriptions are disabled.")

    // If the file is not present, do not try to parse it. This way,
    // Config.EnableTranscription is kept at `false` and the transcription
    // processes are not started.
    return
  }
  err = yaml.Unmarshal(yamlFile, &Config)
  if err != nil {
    log.Fatalf("Unable to read transcription.yaml: %v", err)
  }

  Config.UsedTranscriptionService = getInstanceByServiceName(Config.Service)
}

func getInstanceByServiceName(service string) TranscriptionService {
  s := strings.ToLower(service)

  if strings.EqualFold(s, "aws") || strings.EqualFold(s, "amazon") {
    return GetInstanceOfAwsTranscriptionService()
  }

  if strings.EqualFold(s, "azure") {
    return GetInstanceOfAzureTranscriptionService()
  }

  if strings.EqualFold(s, "gcp") {
    return GetInstanceOfGoogleTranscriptionService()
  }

  log.Fatalf("The transcription service `%s` has not been implemented. Please check the documentation for acceptable values and change the `transcription.yaml` file.", service)
  return nil
}
