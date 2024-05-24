from enum import Enum


class EventType(Enum):
    MediaPause = "MediaPause"
    MediaPlay = "MediaPlay"
    MediaRate = "MediaRate"
    MediaResume = "MediaResume"
    MediaScrobble = "MediaScrobble"
    MediaStop = "MediaStop"
