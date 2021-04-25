package transcription

import tss "github.com/aws/aws-sdk-go/service/transcribestreamingservice"

// Enable or disable transcriptions:
const EnableTranscriptions = true

// Choose here which TranscriptionService to use. Possible function calls are:
// * `GetInstanceOfAwsTranscriptionService()`
// * `GetInstanceOfAzureTranscriptionService()`
// * `GetInstanceOfGoogleTranscriptionService()`
var UsedTranscriptionService = GetInstanceOfGoogleTranscriptionService()
const pcmChunkLen = 1024

// Configure the usage of AWS here
const awsLanguage = tss.LanguageCodeEnUs


// This files contains a bunch of secrets that should *not* appear on streams or
// publicly. Ideally, the secrets will be moved to userspace with the config
// package, but that's for future me.
// Fill the data from the Azure Portal here:
const azureSubscription = "a866dfa9022c4e52994ec9e67ac46be9"
const azureRegion = "westeurope"
const azureLanguage = "en-GB"

// Configure the usage of Google here
const gcpLanguage = "en-UK"
