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
      currentSubtitle: "", //"<p>This is the place for some very long subtitles that might span one or two or even more lines which is totally fine because we thought about that during development.</p>",
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
      timeSinceBegin
    } = message;

    if (type != 'SUBTITLE') {
      return;
    }

    this.setState({currentSubtitle: body});
  }

  render() {
    return html`<div id="subtitles-box" class="flex bg-black text-white text-3xl p-4">
      <span class="w-8/12" dangerouslySetInnerHTML=${{ __html: this.state.currentSubtitle }}></span>
    </div>`
  }
}
