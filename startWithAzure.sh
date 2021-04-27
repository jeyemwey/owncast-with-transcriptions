# Azure
export SPEECHSDK_ROOT="$HOME/speechsdk"
export CGO_CFLAGS="-I$SPEECHSDK_ROOT/include/c_api"
export CGO_LDFLAGS="-L$SPEECHSDK_ROOT/lib/x64 -lMicrosoft.CognitiveServices.Speech.core"
export LD_LIBRARY_PATH="$SPEECHSDK_ROOT/lib/x64:$LD_LIBRARY_PATH"

# Google Cloud
export GOOGLE_APPLICATION_CREDENTIALS="/home/jannik/.ssh/bsc-thesis-308620-fb8eca85c385.json"

go run -tags websockets main.go pkged.go
