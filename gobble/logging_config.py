import logging
from copy import copy
from typing import Any


class CustomFormatter(logging.Formatter):
    def formatMessage(self, record: logging.LogRecord) -> str:
        record_copy = copy(record)
        level_name = record_copy.levelname
        seperator = " " * (8 - len(record_copy.levelname))
        record_copy.__dict__["levelprefix"] = level_name + ":" + seperator
        return super().formatMessage(record_copy)


LOGGING_CONFIG: dict[str, Any] = {
    "version": 1,
    "disable_existing_loggers": False,
    "formatters": {
        "default": {
            "()": "gobble.logging_config.CustomFormatter",
            "fmt": "%(levelprefix)s %(message)s",
        },
    },
    "handlers": {
        "default": {
            "formatter": "default",
            "class": "logging.StreamHandler",
            "stream": "ext://sys.stderr",
        },
    },
    "loggers": {
        "gobble": {"handlers": ["default"], "level": "INFO", "propagate": False},
    },
}
