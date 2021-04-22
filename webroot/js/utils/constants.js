// misc constants used throughout the app

// export const URL_STATUS = `https://watch.owncast.online/api/status`;
export const URL_STATUS = `/api/status`;
export const URL_CHAT_HISTORY = `/api/chat`;
export const URL_CUSTOM_EMOJIS = `/api/emoji`;
export const URL_CONFIG = `/api/config`;

// TODO: This directory is customizable in the config.  So we should expose this via the config API.
// export const URL_STREAM = `https://devstreaming-cdn.apple.com/videos/streaming/examples/bipbop_16x9/bipbop_16x9_variant.m3u8`;
// export const URL_STREAM = `https://watch.owncast.online/hls/stream.m3u8`;
export const URL_STREAM = `/hls/stream.m3u8`;
export const URL_WEBSOCKET = `${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/entry`;

export const TIMER_STATUS_UPDATE = 5000; // ms
export const TIMER_DISABLE_CHAT_AFTER_OFFLINE = 5 * 60 * 1000; // 5 mins
export const TIMER_STREAM_DURATION_COUNTER = 1000;
export const TEMP_IMAGE = 'data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7';

export const OWNCAST_LOGO_LOCAL = '/img/logo.svg';

export const MESSAGE_OFFLINE = 'Stream is offline.';
export const MESSAGE_ONLINE = 'Stream is online.';

export const URL_OWNCAST = 'https://owncast.online'; // used in footer
export const PLAYER_VOLUME = 'owncast_volume';


export const KEY_USERNAME = 'owncast_username';
export const KEY_CUSTOM_USERNAME_SET = 'owncast_custom_username_set'
export const KEY_CHAT_DISPLAYED = 'owncast_chat';
export const KEY_CHAT_FIRST_MESSAGE_SENT = 'owncast_first_message_sent';
export const CHAT_INITIAL_PLACEHOLDER_TEXT = 'Type here to chat, no account necessary.';
export const CHAT_PLACEHOLDER_TEXT = 'Message';
export const CHAT_PLACEHOLDER_OFFLINE = 'Chat is offline.';
export const CHAT_MAX_MESSAGE_LENGTH = 500;
export const CHAT_CHAR_COUNT_BUFFER = 20;
export const CHAT_OK_KEYCODES = [
  'ArrowLeft',
  'ArrowUp',
  'ArrowRight',
  'ArrowDown',
  'Shift',
  'Meta',
  'Alt',
  'Delete',
  'Backspace',
];
export const CHAT_KEY_MODIFIERS = [
  'Control',
  'Shift',
  'Meta',
  'Alt',
];
export const MESSAGE_JUMPTOBOTTOM_BUFFER = 300;

// app styling
export const WIDTH_SINGLE_COL = 730;
export const HEIGHT_SHORT_WIDE = 500;
export const ORIENTATION_PORTRAIT = 'portrait';
export const ORIENTATION_LANDSCAPE = 'landscape';
