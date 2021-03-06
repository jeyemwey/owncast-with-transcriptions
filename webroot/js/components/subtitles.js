import { KEY_CUSTOM_USERNAME_SET } from "../utils/constants.js";
import { getLocalStorage } from "../utils/helpers.js";
import { CALLBACKS } from "../utils/websocket.js";
import htm from '/js/web_modules/htm.js';
import { Component, h } from '/js/web_modules/preact.js';
const html = htm.bind(h);

export default class Subtitles extends Component {
  constructor(props, context) {
    super(props, context);

    this.state = {
      currentSubtitle: "",//"This is the place for some very long subtitles that might span one or two or even more lines which is totally fine because we thought about that during development.",
      webSocketConnected: true,
      queuedSubtitles: [], // type: { body: string, timeSinceBegin: number }
    };
    this.websocket = null;
    this.websocketConnected = this.websocketConnected.bind(this);
    this.websocketDisconnected = this.websocketDisconnected.bind(this);
    this.receivedWebsocketMessage = this.receivedWebsocketMessage.bind(this);
  }

  componentDidMount() {
    this.setupWebSocketCallbacks();
  }

  setupWebSocketCallbacks() {
    this.websocket = this.props.websocket;
    if (this.websocket) {
      this.websocket.addListener(CALLBACKS.RAW_WEBSOCKET_MESSAGE_RECEIVED, this.receivedWebsocketMessage);
      this.websocket.addListener(CALLBACKS.WEBSOCKET_CONNECTED, this.websocketConnected);
      this.websocket.addListener(CALLBACKS.WEBSOCKET_DISCONNECTED, this.websocketDisconnected);
    }
  }

  websocketConnected() {
    this.setState({
      webSocketConnected: true,
    });

    const hasPreviouslySetCustomUsername = getLocalStorage(KEY_CUSTOM_USERNAME_SET);
    if (hasPreviouslySetCustomUsername && !this.props.ignoreClient) {
      this.sendJoinedMessage();
    }
  }

  websocketDisconnected() {
    this.setState({
      webSocketConnected: false,
    });
  }

  receivedWebsocketMessage(message) {
    const {
      body,
      type,
      timeSinceBegin,
      activeHlsSegments
    } = message;

    if (!this.props.isPlaying) {
      return;
    }

    if (type != 'SUBTITLE') {
      return;
    }

    const msgSegmentComparedToGlobalStateSegment = typeof window.activeHlsStream == "undefined" ? 0 : activeHlsSegments[window.activeHlsStream].compareLocale(window.loadedHlsSegment);

    switch (msgSegmentComparedToGlobalStateSegment) {
      case (0): // message and window is equal = display now!
        this.setState({currentSubtitle: body.replace(/<[^>]*>/g, '')});
        break;
      case (-1): // playback state is before than message = come back later
        setTimeout(() => {
          this.receivedWebsocketMessage(message);
        }, 200);
        break;
      case (1): // playback state is later than the message.
        return;
    }
  }

  render() {
    if (!this.props.hasTranscriptionEnabled || !this.props.showSubtitles) {
      return null;
    }

    return html`<div id="subtitles-box" class="absolute z-50 left-0 p-4 w-8/12" style="bottom: 10%; pointer-events: none">
      <span class="bg-black text-white text-3xl leading-tight" dangerouslySetInnerHTML=${{ __html: this.state.currentSubtitle }}></span>
    </div>`
  }
}
