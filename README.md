# Owncast fork for Jannik Volkland's bachelor thesis

This is a fork of the Owncast project which aims to implement an online transcription service into Owncast.
As part of my bachelor thesis, I cannot accept pull requests at this point.
Also, please do not expect the master branch to be as reliable as on the Owncast `master` branch.

This fork appeared after the [v0.0.6 release](https://owncast.online/releases/owncast-0.0.6/).

For any issues besides the transcription, please visit the [Owncast issues page](https://github.com/owncast/owncast/issues) and our [RocketChat](https://owncast.rocket.chat).

## Building from Source with AWS integration

1. Create an AWS account if you don't have one already.
1. In IAM, create an user with "Programmatic Access" and access to `AmazonTranscribeFullAccess`.
1. Generate a key pair and save it in `$HOME/.aws/credentials`:
  ```
  [default]
  aws_access_key_id = <Access key ID>
  aws_secret_access_key = <Secret access key>
  ```
1. Configure the application acc. to the instructions below.
1. Start the application with `./start.sh`.
  
## Building from Source with Azure integration

1. Ensure you have the gcc compiler configured.
1. Install the [Go toolchain](https://golang.org/dl/).
1. Install the [Microsoft Azure Speech SDK for Go](https://docs.microsoft.com/en-us/azure/cognitive-services/speech-service/quickstarts/setup-platform?tabs=dotnet%2Clinux%2Cjre%2Cbrowser&pivots=programming-language-go)
1. Clone the repo.  `git clone https://github.com/owncast/owncast`
1. Update `start.sh` to contain the right paths.
1. `./start.sh` will run from source.
1. Point your [broadcasting software](https://owncast.online/docs/broadcasting/) at your new server and start streaming.

## Building from Source with Google integration

1. Make sure you have a GCP account with a project and a service account. Activate the "Speech-to-Text" feature and download a JSON key.
1. Clone the repo (see Azure).
1. Edit `start.sh` to contain the path to your key file.
1. `./start.sh`

## Configuration

Copy the `transcription-example.yaml` file to `transcription.yaml` and edit accordingly.
Edit the configuration according to the documentation.

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

Project chat: [Join us on Rocket.Chat](https://owncast.rocket.chat/home) if you want to contribute, follow along, or if you have questions.

Jannik Volkland - [@jannik@uelfte.club](https://uelfte.club/@jannik) - email [jvolklan@th-koeln.de](mailto:jvolklan@th-koeln.de)

Project Link: [https://github.com/owncast/owncast](https://github.com/owncast/owncast)
