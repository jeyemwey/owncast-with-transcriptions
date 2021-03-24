# Owncast fork for Jannik Volkland's bachelor thesis

This is a fork of the Owncast project which aims to implement an online transcription service into Owncast.
As part of my bachelor thesis, I cannot accept pull requests at this point.
Also, please do not expect the master branch to be as reliable as on the Owncast `master` branch.

This fork appeared after the [v0.0.6 release](https://owncast.online/releases/owncast-0.0.6/).

For any issues besides the transcription, please visit the [Owncast issues page](https://github.com/owncast/owncast/issues) and our [RocketChat](https://owncast.rocket.chat).

## Building from Source with Google integration

1. Make sure you have a GCP account with a project and a service account. Activate the "Speech-to-Text" feature and download a JSON key.
2. Clone the repo.
3. Edit `startWithGoogle.sh` to contain the path to your key file.
4. `./startWithGoogle.sh`

## Building from Source with Azure integration

1. Ensure you have the gcc compiler configured.
1. Install the [Go toolchain](https://golang.org/dl/).
1. Install the [Microsoft Azure Speech SDK for Go](https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/quickstarts/setup-platform?tabs=dotnet%2Clinux%2Cjre%2Cbrowser&pivots=programming-language-go)
1. Clone the repo.  `git clone https://github.com/owncast/owncast`
1. `./startWithAzure.sh` will run from source.
1. Point your [broadcasting software](https://owncast.online/docs/broadcasting/) at your new server and start streaming.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

Project chat: [Join us on Rocket.Chat](https://owncast.rocket.chat/home) if you want to contribute, follow along, or if you have questions.

Jannik Volkland - [@jannik@uelfte.club](https://uelfte.club/@jannik) - email [jvolklan@th-koeln.de](mailto:jvolklan@th-koeln.de)

Project Link: [https://github.com/owncast/owncast](https://github.com/owncast/owncast)
