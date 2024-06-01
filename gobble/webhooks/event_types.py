from enum import Enum


class EventType(Enum):
    MediaPause = "media_pause"
    MediaPlay = "media_play"
    MediaRate = "media_rate"
    MediaResume = "media_resume"
    MediaScrobble = "media_scrobble"
    MediaStop = "media_stop"
