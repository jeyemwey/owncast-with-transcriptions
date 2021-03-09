# Owncast fork for Jannik Volkland's bachelor thesis

This is a fork of the Owncast project which aims to implement an online transcription service into Owncast.
As part of my bachelor thesis, I cannot accept pull requests at this point.
Also, please do not expect the master branch to be as reliable as on the Owncast `master` branch.

This fork appeared after the [v0.0.6 release](https://owncast.online/releases/owncast-0.0.6/).

For any issues besides the transcription, please visit the [Owncast issues page](https://github.com/owncast/owncast/issues) and our [RocketChat](https://owncast.rocket.chat).

## Building from Source

1. Ensure you have the gcc compiler configured.
1. Install the [Go toolchain](https://golang.org/dl/).
1. Clone the repo.  `git clone https://github.com/owncast/owncast`
1. `go run main.go pkged.go` will run from source.
1. Point your [broadcasting software](https://owncast.online/docs/broadcasting/) at your new server and start streaming.

There is also a supplied `Dockerfile` so you can spin it up from source with little effort.  [Read more about running from source](https://owncast.online/docs/building/).

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

Project chat: [Join us on Rocket.Chat](https://owncast.rocket.chat/home) if you want to contribute, follow along, or if you have questions.

Jannik Volkland - [@jannik@uelfte.club](https://uelfte.club/@jannik) - email [jvolklan@th-koeln.de](mailto:jvolklan@th-koeln.de)

Project Link: [https://github.com/owncast/owncast](https://github.com/owncast/owncast)
