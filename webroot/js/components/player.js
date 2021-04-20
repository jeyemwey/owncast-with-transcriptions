// https://docs.videojs.com/player

import { PLAYER_VOLUME, URL_STREAM } from '../utils/constants.js';
import { getLocalStorage, setLocalStorage } from '../utils/helpers.js';
import '/js/web_modules/@videojs/http-streaming/dist/videojs-http-streaming.min.js';
import videojs from '/js/web_modules/videojs/core.js';

const VIDEO_ID = 'video';

// Video setup
const VIDEO_SRC = {
  src: URL_STREAM,
  type: 'application/x-mpegURL',
};
const VIDEO_OPTIONS = {
  autoplay: false,
  liveui: true,
  preload: 'auto',
  controlBar: {
    progressControl: {
      seekBar: false,
    },
  },
  html5: {
    vhs: {
      // used to select the lowest bitrate playlist initially. This helps to decrease playback start time. This setting is false by default.
      enableLowInitialPlaylist: true,
    },
  },
  liveTracker: {
    trackingThreshold: 0,
  },
  sources: [VIDEO_SRC],
};

export const POSTER_DEFAULT = `/img/logo.png`;
export const POSTER_THUMB = `/thumbnail.jpg`;

class OwncastPlayer {
  constructor() {
    window.VIDEOJS_NO_DYNAMIC_STYLE = true; // style override

    this.vjsPlayer = null;

    this.appPlayerReadyCallback = null;
    this.appPlayerPlayingCallback = null;
    this.appPlayerPauseCallback = null;
    this.appPlayerEndedCallback = null;

    // bind all the things because safari
    this.startPlayer = this.startPlayer.bind(this);
    this.handleReady = this.handleReady.bind(this);
    this.handlePlaying = this.handlePlaying.bind(this);
    this.handlePause = this.handlePause.bind(this);
    this.handleVolume = this.handleVolume.bind(this);
    this.handleEnded = this.handleEnded.bind(this);
    this.handleError = this.handleError.bind(this);
  }

  init() {
    videojs.Vhs.xhr.beforeRequest = options => {
      if (options.uri.match('m3u8')) {
        const cachebuster = Math.round(new Date().getTime() / 1000);
        options.uri = `${options.uri}?cachebust=${cachebuster}`;
      }
      return options;
    };

    this.vjsPlayer = videojs(VIDEO_ID, VIDEO_OPTIONS);

    window.vjsPlayer = this.vjsPlayer;
    window.videojs = this.vjsPlayer;

    this.addAirplay();
    this.vjsPlayer.ready(this.handleReady);

    this.vjsPlayer.overlay({
      content: 'Default overlay content',
      debug: true,
      overlays: [{
        content: 'This is a test-overlay!',
        start: 'play',
        end: 'pause'
      }]
    });
  }

  setupPlayerCallbacks(callbacks) {
    const { onReady, onPlaying, onPause, onEnded, onError } = callbacks;

    this.appPlayerReadyCallback = onReady;
    this.appPlayerPlayingCallback = onPlaying;
    this.appPlayerPauseCallback = onPause;
    this.appPlayerEndedCallback = onEnded;
    this.appPlayerErrorCallback = onError;
  }

  // play
  startPlayer() {
    this.log('Start playing');
    const source = { ...VIDEO_SRC };

    this.vjsPlayer.volume(getLocalStorage(PLAYER_VOLUME) || 1);
    this.vjsPlayer.src(source);
    // this.vjsPlayer.play();
  }

  handleReady() {
    this.log('on Ready');
    this.vjsPlayer.on('error', this.handleError);
    this.vjsPlayer.on('playing', this.handlePlaying);
    this.vjsPlayer.on('paused', this.handlePause);
    this.vjsPlayer.on('volumechange', this.handleVolume);
    this.vjsPlayer.on('ended', this.handleEnded);
    this.vjsPlayer.on('loadedmetadata', (ev) => console.log(ev));

    if (this.appPlayerReadyCallback) {
      // start polling
      this.appPlayerReadyCallback();
    }
  }

  handleVolume() {
    setLocalStorage(PLAYER_VOLUME, this.vjsPlayer.muted() ? 0 : this.vjsPlayer.volume());
  }

  handlePlaying() {
    console.log('on Playing');
    if (this.appPlayerPlayingCallback) {
      // start polling
      this.appPlayerPlayingCallback();
    }
  }

  handlePause() {
    console.log('on pause');
    if (this.appPlayerPauseCallback) {
      // start polling
      this.appPlayerPauseCallback();
    }
  }

  handleEnded() {
    this.log('on Ended');
    if (this.appPlayerEndedCallback) {
      this.appPlayerEndedCallback();
    }
  }

  handleError(e) {
    this.log(`on Error: ${JSON.stringify(e)}`);
    if (this.appPlayerEndedCallback) {
      this.appPlayerEndedCallback();
    }
  }

  log(message) {
    // console.log(`>>> Player: ${message}`);
  }

  addAirplay() {
    videojs.hookOnce('setup', function (player) {
      if (window.WebKitPlaybackTargetAvailabilityEvent) {
        var videoJsButtonClass = videojs.getComponent('Button');
        var concreteButtonClass = videojs.extend(videoJsButtonClass, {
          // The `init()` method will also work for constructor logic here, but it is
          // deprecated. If you provide an `init()` method, it will override the
          // `constructor()` method!
          constructor: function () {
            videoJsButtonClass.call(this, player);
          },

          handleClick: function () {
            const videoElement = document.getElementsByTagName('video')[0];
            videoElement.webkitShowPlaybackTargetPicker();
          },
        });

        var concreteButtonInstance = this.vjsPlayer.controlBar.addChild(
          new concreteButtonClass()
        );
        concreteButtonInstance.addClass('vjs-airplay');
      }
    });
  }
}

export { OwncastPlayer };
