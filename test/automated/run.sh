#!/bin/bash

TEMP_DB=$(mktemp)

# Install the node test framework
npm install --silent > /dev/null

# Download a specific version of ffmpeg
if [ ! -d "ffmpeg" ]; then
  mkdir ffmpeg
  pushd ffmpeg > /dev/null
  curl -sL https://github.com/vot/ffbinaries-prebuilt/releases/download/v4.2.1/ffmpeg-4.2.1-linux-64.zip --output ffmpeg.zip > /dev/null
  unzip -o ffmpeg.zip > /dev/null
  PATH=$PATH:$(pwd)
  popd > /dev/null
fi

pushd ../.. > /dev/null

## ADD AZURE SPEECH SDK
sudo apt-get update
sudo apt-get install build-essential libssl1.0.0 libasound2 wget
export SPEECHSDK_ROOT="$HOME/speechsdk"
mkdir -p "$SPEECHSDK_ROOT"
wget -O SpeechSDK-Linux.tar.gz https://aka.ms/csspeech/linuxbinary
tar --strip 1 -xzf SpeechSDK-Linux.tar.gz -C "$SPEECHSDK_ROOT"
ls -l "$SPEECHSDK_ROOT"
export CGO_CFLAGS="-I$SPEECHSDK_ROOT/include/c_api"
export CGO_LDFLAGS="-L$SPEECHSDK_ROOT/lib/<architecture> -lMicrosoft.CognitiveServices.Speech.core"
export LD_LIBRARY_PATH="$SPEECHSDK_ROOT/lib/<architecture>:$LD_LIBRARY_PATH"
      
# Build and run owncast from source
go build -o owncast main.go pkged.go
./owncast -database $TEMP_DB &
SERVER_PID=$!

popd > /dev/null
sleep 5

# Start streaming the test file over RTMP to
# the local owncast instance.
ffmpeg -hide_banner -loglevel panic -stream_loop -1 -re -i test.mp4 -vcodec libx264 -profile:v main -sc_threshold 0 -b:v 1300k -acodec copy -f flv rtmp://127.0.0.1/live/abc123 &
FFMPEG_PID=$!

function finish {
  rm $TEMP_DB
  kill $SERVER_PID $FFMPEG_PID
}
trap finish EXIT

echo "Waiting..."
sleep 13

# Run the tests against the instance.
npm test